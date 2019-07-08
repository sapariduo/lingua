// +build !stanfordrel

package treebank

import "github.com/sapariduo/lingua"

var dependencyTable map[string]lingua.DependencyType = map[string]lingua.DependencyType{
	"dep":          lingua.Dep,
	"root":         lingua.Root,
	"nsubj":        lingua.NSubj,
	"nsubjpass":    lingua.NSubjPass,
	"dobj":         lingua.DObj,
	"iobj":         lingua.IObj,
	"csubj":        lingua.CSubj,
	"csubjpass":    lingua.CSubjPass,
	"ccomp":        lingua.CComp,
	"xcomp":        lingua.XComp,
	"nummod":       lingua.NumMod,
	"appos":        lingua.Appos,
	"nmod":         lingua.NMod,
	"acl":          lingua.ACl,
	"acl:relcl":    lingua.ACl_RelCl,
	"det":          lingua.Det,
	"det:predet":   lingua.Det_PreDet,
	"amod":         lingua.AMod,
	"neg":          lingua.Neg,
	"case":         lingua.Case,
	"nmod:npmod":   lingua.NMod_NPMod,
	"nmod:tmod":    lingua.NMod_TMod,
	"nmod:poss":    lingua.NMod_Poss,
	"advcl":        lingua.AdvCl,
	"advmod":       lingua.AdvMod,
	"compound":     lingua.Compound,
	"compound:prt": lingua.Compound_Part,
	"name":         lingua.Name,
	"mwe":          lingua.MWE,
	"foreign":      lingua.Foreign,
	"goeswith":     lingua.GoesWith,
	"list":         lingua.List,
	"dislocated":   lingua.Dislocated,
	"parataxis":    lingua.Parataxis,
	"remnant":      lingua.Remnant,
	"reparandum":   lingua.Reparandum,
	"vocative":     lingua.Vocative,
	"discourse":    lingua.Discourse,
	"expl":         lingua.Expl,
	"aux":          lingua.Aux,
	"auxpass":      lingua.AuxPass,
	"cop":          lingua.Cop,
	"mark":         lingua.Mark,
	"punct":        lingua.Punct,
	"conj":         lingua.Conj,
	"cc":           lingua.Coordination,
	"cc:preconj":   lingua.CC_PreConj, // https://github.com/UniversalDependencies/docs/issues/221
	"conj:preconj": lingua.CC_PreConj, // https://github.com/UniversalDependencies/docs/issues/221

	"-NULL-": lingua.NoDepType,
}
