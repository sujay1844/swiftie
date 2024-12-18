{
  description = "DNS server to serve Taylor Swift lyrics";

  inputs = { nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable"; };

  outputs = { self, nixpkgs }:
    let
      system = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages.${system};
    in {
      devShells.${system}.default = pkgs.mkShell {
        buildInputs = with pkgs; [ go golangci-lint gopls air sqlite sqlc ];
        shellHook =
          # bash
          ''
            if [ -f .env ]; then
              eval $(cat .env | sed 's/^/export /')
              echo "Environment variables sourced."
            fi
          '';
      };
    };
}
