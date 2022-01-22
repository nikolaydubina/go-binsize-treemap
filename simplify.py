import sys

values['name'] = '>'
def flatten(d):
    vals = []
    for k, v in d.items():
        v['name'] = k
        vals.append(v)
    return vals

def transform(d):
    # transform [A, B., X] into  [A, B/, _self_, X]  if [A, B/] also exists already.
    maybecopy = None
    for k in d:
        if k.endswith('.'):
            dirkey = k[:-1] + '/'
            if dirkey in d:
                if maybecopy is None:
                    maybecopy = d.copy()
                maybecopy[dirkey]['children']['· self'] = d[k]
                del(maybecopy[k])
    if maybecopy is not None:
        d = maybecopy

    for k, v in list(d.items()):
        if type(v) == type({}):
            ename, v = transform(v)
            del d[k]
            d[k+ename] = v
    ename = ""
    if "children" in d:
        c = flatten(d['children'])
        if len(c) == 1 and "children" in c[0]:
            c0 = c[0]
            ename = c0["name"]
            c = c0["children"]
        d['children'] = c
    return ename, d
