{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-25.05";
    flake-utils.url = "github:numtide/flake-utils";
  };
  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import nixpkgs {
        inherit system;
      };
      fastlauncher_next-package = pkgs.callPackage ./package.nix {};
    in {
      packages = rec {
        fastlauncher_next = fastlauncher-package;
        default = fastlauncher_next;
      };

      apps = rec {
        fastlauncher_next = flake-utils.lib.mkApp {
          drv = self.packages.${system}.fastlauncher_next;
        };
        default = fastlauncher_next;
      };

      devShells.default = pkgs.mkShell {
        packages = (with pkgs; [
          go
        ]);
      };
    });
}
