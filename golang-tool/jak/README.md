# jak

Just Another http Klient

version: pre-beta

## Usage Examples

* simple
* batch
* chain - under experimental

### Simple

```bash
jak req GET https://example.com

jak req POST https://example.com/api -H "Content-Type: application/json" -j '{"key":"value"}'
```

### Batch

```bash
jak bat your-setting.toml
```

- [sample toml single](test/fixtures/bat_simple.toml)
- [sample tomle multiple](test/fixtures/bat_multiple.toml)

### Chain

variable extraction and substitution

```bash
jak chain your-setting.toml
```

- [sample toml](test/fixtures/chain.toml)

## Installation

```bash
# Clone the repository
git clone https://github.com/ymatsukawa/jak.git
cd jak

# Build the application
make build

# Or install it to your GOPATH
make install
```
