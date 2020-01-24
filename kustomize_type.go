package main

type Kustomization struct {
	Bases                 []string           `yaml:"bases,omitempty"`
	Patches               []string           `yaml:"patches,omitempty"`
	PatchesStrategicMerge []string           `yaml:"patchesStrategicMerge,omitempty"`
	Images                []*Image           `yaml:"images,omitempty"`
	SecretGenerator       []*SecretGenerator `yaml:"secretGenerator,omitempty"`
}

type Image struct {
	Name   string `yaml:"name,omitempty"`
	NewTag string `yaml:"newTag,omitempty"`
}

type SecretGenerator struct {
	Env   string   `yaml:"env,omitempty"`
	Name  string   `yaml:"name,omitempty"`
	Type  string   `yaml:"type,omitempty"`
	Files []string `yaml:"files,omitempty"`
}
