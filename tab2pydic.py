import re
import sys

undefre = re.compile(r'^\s*0\s+U\s+',re.A)
entriesre = re.compile(r'^\s*([0-9a-fA-F]+)\s+([0-9]+)\s+(\S+)\s+(.*)$',re.A)

tmplpart = r'''
<
 (?:
    [^<>]
  | (?:
     <
      (?:
        [^<>]
      | (?:
          <
            (?:
              [^<>]
            | (?:
                <
                  (?:
                    [^<>]
                  | (?:
                      <
                        (?:
                          [^<>]
                        | (?:
                            <
                              (?:
                                [^<>]
                              | (?:
                                  <
                                    (?:
                                      [^<>]
                                    | (?:
                                        <
                                          (?:
                                            [^<>]
                                          | (?:
                                              <
                                                [^<>]*
                                              >
                                            )
                                          )*
                                        >
                                      )
                                    )*
                                  >
                                )
                              )*
                            >
                          )
                        )*
                      >
                    )
                  )*
                >
              )
            )*
          >
        )
      )*
     >
    )
 )*
>
'''
parengroup = tmplpart.replace("<", r'\(').replace(">", r'\)')

cpppath = '''
     (?:
        \([^)]*\)                   # (anonymous namespace)
      | \{[^}]*\}                   # {lambda ...}
      | ~?                          # maybe ~destructor
        (?:
           \$?\w+                      # name
         | operator(?:[^\(]+|\(\))  # special name
        )
        (?:'''+tmplpart+''')?       # maybe <tmpl params>
        (?:'''+parengroup+''')?     # maybe function args
        (?:\sconst)?                # maybe " const"
     )::
  | [a-zA-Z]+_
'''

cpppathre = re.compile(cpppath, re.X|re.A)

cppsymre = re.compile(r'''
^
# prefix
(?:guard\svariable\sfor\s)?
(
  # type
  (?:
    (?:
       \w
     | ::
     | -
     | \*
     | \&
     | (?:'''+tmplpart+''')
    )+
    \s
  )*
)
(
  # path / namespace
  (?:'''+cpppath+''')*
)
(
  \{[^}]*\}                   # {lambda ...}
|
  # object name
  ~?                          # maybe destructor
  (?:
     \$?\w+                      # name
   | operator(?:[^\(]+|\(\))  # special name
   | \._\d+                   # ._NNN
  )
  (?:\[[^][]*\])?              # ABI spec
  (?:'''+tmplpart+''')?       # maybe <tmpl params>
  (?:'''+parengroup+'''
      (?:\sconst)?                # maybe " const"
      (?:\s\[.*\])?               # maybe words
      (?:\s\(
          (?:\.(?:constprop|part|isra).\d+)+
      \))?    # maybe (.constprop.NNN)
  )?
  (?:\*+)?                    # maybe ptr
)
$
''', re.X|re.A)

gopathparts = r'''
    \(
      (?:
         [^()]
       | \( [^()]* \)
      )*
    \)\.                  # (*Foo)., (*struct{... () }).
  | struct\s\{
     (?:
        [^{}]
      | \{[^{}]*\}
     )*
    \}\.                  # struct { ... }.
  | \$?(?:\w|-|%)+\.      # fun.
  | glob\.\.              # glob..
  | \.gobytes\.           # .gobytes.
  | (?:\w|\.|-|%)+/       # path/
'''

gopathpartsre = re.compile(gopathparts, re.X|re.A)

golastpart = r'''
      (?:
        (?:
            \.?
            (?:
               (?:\w|-|%)+    # regular name
             | \( [^()]* \)   # v1.x go 
            )
            (?:-fm)?)  # name
      | struct\s\{
           (?:
              [^{}]
            | \{[^{}]*\}
           )*
          \}                    # struct { ... }
      )
      (?:,
         (?:
          (?:(?:\w|\.|-|%)+/)*
          (?:\w|-|%|\.)+
          | interface\s\{ (?: [^{}] | \{ [^{}]* \} )* \}
         )
      )?    # opt ,xxx interface suffix
    | initdone\.
    | initdone·
'''

golastpartre = re.compile('('+golastpart + ')$', re.X|re.A)

gosymre = re.compile(r'''
^
  (
    (?:go\.itab\.\*?)?
  )
  (
    (?:'''+gopathparts+''')*
  )
  (
    '''+golastpart+'''
  )
$
''', re.X|re.A)

vtables_sz = 0
typeinfo_sz = 0
init_sz = 0
rest_sz = 0
gotyp_sz = 0

values = {}

def store_rec(values, path, sz):
    if len(path) == 0:
        v = values.get("value", 0)
        v += sz
        values["value"] = v
    else:
        n = path[0]
        d1 = values.get("children", {})
        d = d1.get(n, {})
        d1[n] = d
        values["children"] = d1
        store_rec(d, path[1:], sz)

def store(path, sz):
    global values
    store_rec(values, path, sz)

with open(sys.argv[1]) as f:
    for line in f:
        if undefre.match(line) is not None:
            # undefined (external) symbol, skip
            continue
        if line.endswith('\n'):
            line = line[:-1]

        m = entriesre.match(line)
        if m is None:
            print("\nunknown format:", line, file=sys.stderr)
            continue

        typ = m.group(3).strip()
        if typ == 'U':
            # external object, skip
            continue

        sz = m.group(2).strip()
        try:
            sz = int(sz)
        except:
            print("\nunknown size format:", m.group(2), file=sys.stderr)
            continue
        if sz == 0:
            # skip over no-size objects
            continue

        sym = m.group(4).strip()
        if sym == '':
            continue
        if sym.startswith('construction vtable ') or sym.startswith('vtable for '):
            # c++ virtual diamond constructors
            vtables_sz += sz
            continue
        if sym.startswith('__static_initialization_and_destruction'):
            # c++ static initializers
            init_sz += sz
            continue
        if sym.startswith('typeinfo '):
            # c++ type metadata
            typeinfo_sz += sz
            continue
        if sym.startswith('type..'):
            # go generated type equality and hash functions
            gotyp_sz += sz
            continue

        parts = cppsymre.match(sym)
        if parts is not None:
            prefix = ['c/c++ · ']
            partsre = cpppathre

        if parts is None:
            parts = gosymre.match(sym)
            if parts is not None:
                prefix = ['go · ']
                partsre = gopathpartsre

        if parts is None:
            print("\nunknown",typ,"sym format:", sym, file=sys.stderr)
            rest_sz += sz
            continue

        path = parts.group(2)
        path = partsre.findall(path)
        name = parts.group(1) + parts.group(3)

        fullpath = prefix+path+[name]

        store(fullpath, sz)

store(['c/c++ · ','VTABLES'], vtables_sz)
store(['c/c++ · ','TYPEDATA'], typeinfo_sz)
store(['c/c++ · ','INITIALIZERS'], init_sz)
store(['go · ','TYPEDATA'], gotyp_sz)
store(['UNKNOWN'], rest_sz)
