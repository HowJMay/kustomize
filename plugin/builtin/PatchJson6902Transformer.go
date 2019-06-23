// Code generated by pluginator on PatchJson6902Transformer; DO NOT EDIT.
package builtin

import (
	"sigs.k8s.io/kustomize/v3/pkg/ifc"
	"sigs.k8s.io/kustomize/v3/pkg/patch/transformer"
	"sigs.k8s.io/kustomize/v3/pkg/resmap"
	"sigs.k8s.io/kustomize/v3/pkg/types"
	"sigs.k8s.io/yaml"
)

type PatchJson6902TransformerPlugin struct {
	ldr     ifc.Loader
	Patches []types.PatchJson6902 `json:"patches,omitempty" yaml:"patches,omitempty"`
}

//noinspection GoUnusedGlobalVariable
func NewPatchJson6902TransformerPlugin() *PatchJson6902TransformerPlugin {
	return &PatchJson6902TransformerPlugin{}
}

func (p *PatchJson6902TransformerPlugin) Config(
	ldr ifc.Loader, rf *resmap.Factory, c []byte) (err error) {
	p.ldr = ldr
	p.Patches = nil
	return yaml.Unmarshal(c, p)
}

func (p *PatchJson6902TransformerPlugin) Transform(m resmap.ResMap) error {
	t, err := transformer.NewPatchJson6902Factory(p.ldr).
		MakePatchJson6902Transformer(p.Patches)
	if err != nil {
		return err
	}
	return t.Transform(m)
}
