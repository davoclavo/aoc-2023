top@{ ... }: {
  imports = [ ./flix ];

  perSystem = { pkgs, config, ... }: {
    devenv.shells.default = {
      # See https://devenv.sh/packages/
      packages = with pkgs; [
        jq
        aoc-cli
        watchexec
        just
      ];

      # See https://devenv.sh/languages/
      languages.go.enable = true;
      languages.scala.enable = true;
      # languages.ocaml.enable = true;

      languages.java.jdk.package = pkgs.graalvm-ce;
      # languages.ocaml.packages = pkgs.ocaml-ng.ocamlPackages_5_1;
    };

  };
}
