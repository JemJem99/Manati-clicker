components {
  id: "Acuario"
  component: "/main/Acuario.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"aquarium\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/imatge.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.2
    y: 0.2
    z: 0.2
  }
}
