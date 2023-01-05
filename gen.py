

import os
import sys
from string import Template


def ps_list(n, prefix, suffix):
	l = [0] * n
	for i in range(len(l)):
		l[i] = "{}{}{}".format(prefix, i+1, suffix)
	return ", ".join(l)


def tany_list(n):
	return ps_list(n, "T", " any")

def aany_list(n):
	return ps_list(n, "A", " any")

def tonly_list(n):
	return ps_list(n, "T", "")

def vonly_list(n):
	return ps_list(n, "v", "")

def return_field_list(n):
	return ps_list(n, "t.V", "")

def pms_list(n, prefix, middle, suffix, j = ", "):
	l = [0] * n
	for i in range(len(l)):
		l[i] = "{}{}{}{}{}".format(prefix, i+1, middle, i+1, suffix)
	return j.join(l)

def argt_list(n):
	return pms_list(n, "v", " T", "")


def line_list(n, prefix, middle, suffix):
	ret = ""
	for i in range(n):
		ret += "{}{}{}{}{}\n".format(prefix, i+1, middle, i+1, suffix)
	return ret

def field_list(n):
	return pms_list(n, "\tV", " T", "", "\n")

def set_field_list(n):
	return pms_list(n, "\tt.V", " = v", "", "\n")

def cloure(s):
	if s:
		return "(" + s + ")"
	return ""

def concat(a, b):
	if a and b:
		return a + ", " + b
	return a or b


def generate_retprocx(narg, nret):

	header = """package myhig
"""
	fromTemplate = Template("""
func GetProcOfFunc${arg}r${ret}[${artany_list}](_ func(${args_list}) (${tonly_list})) (*RetProc${ret}[${tonly_list}]) {
	rp := &RetProc${ret}[${tonly_list}]{}
	return rp
}
""")


	# with sys.stdout as f:
	with open("gen_retproc.go", "w") as f:
		f.write(header)
		for arg in range(narg+1):
			for ret in range(1, nret+1):
				d = {
						'arg': arg,
						'ret': ret,
						"artany_list": concat(aany_list(arg), tany_list(ret)),
						"args_list": ps_list(arg, "_ A", ""),
						"tonly_list": tonly_list(ret),
						}
				f.write(fromTemplate.substitute(d))


def generate_tuplex(nt):

	header = """package myhig
"""

	fromTemplate = Template("""
type Tuple${n}[${tany_list}]struct {
${field_list}
}

func Tw${n}[${tany_list}](${argt_list}) Tuple${n}[${tonly_list}] {
	return Tuple${n}[${tonly_list}]{${vonly_list}}
}

func (t *Tuple${n}[${tonly_list}]) Unwrap() (${tonly_list}) {
	return ${return_field_list}
}

func (t *Tuple${n}[${tonly_list}]) Set(${argt_list}) {
${set_field_list}
}

func (t *Tuple${n}[${tonly_list}]) setFail(err error) {
	setError(&${last_value}, err)
}

type RetProc${n}[${tany_list}]struct {
	Tuple${n}[${tonly_list}]
	Proc
}

func NewRetProc${n}[${tany_list}]() *RetProc${n}[${tonly_list}] {
	return &RetProc${n}[${tonly_list}]{}
}

func (rt *RetProc${n}[${tonly_list}]) ErrorReturn(err error) {
	if err != nil {
		rt.Tuple${n}.setFail(err)
		rt.ReturnOnly()
	}
}

func (rt *RetProc${n}[${tonly_list}]) FalseReturn(b bool) {
	if !b {
		rt.Tuple${n}.setFail(ErrIsFalse)
		rt.ReturnOnly()
	}
}

func (rt *RetProc${n}[${tonly_list}]) Return(${argt_list}) {
	rt.Tuple${n}.Set(${vonly_list})
	rt.ReturnOnly()
}

func (rt *RetProc${n}[${tonly_list}]) IfReturn(b bool, ${argt_list}) {
	if b {
		rt.Tuple${n}.Set(${vonly_list})
		rt.ReturnOnly()
	}
}

func (rt *RetProc${n}[${tonly_list}]) Dov(fn func()) *Tuple${n}[${tonly_list}] {
	rt.Do(fn)
	return &rt.Tuple${n}
}

func (rt *RetProc${n}[${tonly_list}]) Dow(fn func()) (${tonly_list}) {
	rt.Do(fn)
	return rt.Tuple${n}.Unwrap()
}

type RetTuple${n}[${tany_list}]struct {
	Tuple${n}[${tonly_list}]
}

func Must${n}[${tany_list}](${argt_list}) *RetTuple${n}[${tonly_list}] {
	return &RetTuple${n}[${tonly_list}]{Tuple${n}[${tonly_list}]{${vonly_list}}}
}

func (t *RetTuple${n}[${tonly_list}]) OrReturnTo(ra failReturnAble) (${tonly_prefix_list}) {
	tupleMust(ra, ${last_value}, nil)
	return ${return_prefix_field_list}
}

func (t *RetTuple${n}[${tonly_list}]) OrDoReturnTo(fn func(error), ra failReturnAble) (${tonly_prefix_list}) {
	tupleMust(ra, ${last_value}, fn)
	return ${return_prefix_field_list}
}

func (t *RetTuple${n}[${tonly_list}]) OrFunc(fn func() ${closure_tonly_prefix_list}) (${tonly_prefix_list}) {
	if isFail(${last_value}) {
		${return_fn}
	}
	return ${return_prefix_field_list}
}

func (t *RetTuple${n}[${tonly_list}]) Or(${argt_prefix_list}) (${tonly_prefix_list}) {
	if isFail(${last_value}) {
		return ${vonly_prefix_list}
	}
	return ${return_prefix_field_list}
}

""")


	# with sys.stdout as f:
	with open("gen_tuple.go", "w") as f:
		f.write(header)
		for n in range(1, nt+1):
			d = {
					'n': n,
					"tany_list": tany_list(n),
					"tonly_list": tonly_list(n),
					"argt_list": argt_list(n),
					"vonly_list": vonly_list(n),
					"field_list": field_list(n),
					"return_field_list": return_field_list(n),
					"set_field_list": set_field_list(n),
					"last_value": "t.V{}".format(n),
					"return_prefix_field_list": return_field_list(n-1),
					"argt_prefix_list": argt_list(n-1),
					"vonly_prefix_list": vonly_list(n-1),
					"tonly_prefix_list": tonly_list(n-1),
					"closure_tonly_prefix_list": cloure(tonly_list(n-1)),
					"return_fn": "return fn()" if n-1 > 0 else "return",
					}
			f.write(fromTemplate.substitute(d))




generate_retprocx(10, 10)
generate_tuplex(10)

