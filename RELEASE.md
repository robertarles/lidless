# Release Process

This document describes how to create a new release of Lidless.

## Prerequisites

- Write access to the repository
- All changes committed and pushed to `main`
- Tests passing
- Version number decided (follow [Semantic Versioning](https://semver.org/))

## Release Steps

### 1. Update Version References (if applicable)

Update any version strings in documentation or code if needed.

### 2. Commit Changes

```bash
git add .
git commit -m "Prepare vX.Y.Z release"
git push origin main
```

### 3. Create and Push Tag

```bash
# Create an annotated tag
git tag -a vX.Y.Z -m "Release vX.Y.Z"

# Push the tag to trigger the release workflow
git push origin vX.Y.Z
```

Replace `X.Y.Z` with your version number (e.g., `v1.0.0`, `v1.2.3`).

### 4. Monitor GitHub Actions

1. Go to https://github.com/robertarles/lidless/actions
2. Watch the "Release" workflow execute
3. Verify it completes successfully

### 5. Verify the Release

1. Go to https://github.com/robertarles/lidless/releases
2. Verify the new release appears with:
   - Correct version number
   - Both binary archives (amd64 and arm64)
   - Installation instructions in the release notes

### 6. Test the Release

Download and test both binaries:

```bash
# Test Intel binary
curl -L https://github.com/robertarles/lidless/releases/download/vX.Y.Z/lidless-vX.Y.Z-darwin-amd64.tar.gz | tar xz
./lidless-darwin-amd64

# Test Apple Silicon binary
curl -L https://github.com/robertarles/lidless/releases/download/vX.Y.Z/lidless-vX.Y.Z-darwin-arm64.tar.gz | tar xz
./lidless-darwin-arm64
```

## Release Workflow Details

The release is automated via GitHub Actions (`.github/workflows/release.yml`):

1. Triggered by pushing a tag matching `v*`
2. Builds binaries for:
   - macOS Intel (darwin/amd64)
   - macOS Apple Silicon (darwin/arm64)
3. Packages each binary as a `.tar.gz` archive
4. Creates a GitHub release with:
   - The tagged version
   - Both binary archives attached
   - Installation instructions in the release body

## Versioning Guidelines

Follow [Semantic Versioning](https://semver.org/):

- **Major (vX.0.0)**: Breaking changes, incompatible API changes
- **Minor (v1.X.0)**: New features, backwards-compatible
- **Patch (v1.0.X)**: Bug fixes, backwards-compatible

Examples:
- `v0.1.0` - Initial development release
- `v1.0.0` - First stable release
- `v1.1.0` - Added new feature
- `v1.1.1` - Fixed bug in v1.1.0

## Troubleshooting

### Release workflow fails

1. Check the Actions logs at https://github.com/robertarles/lidless/actions
2. Common issues:
   - Build errors: Fix code and create a new tag
   - Permission errors: Verify repository settings allow Actions to create releases

### Need to update a release

1. Delete the tag locally and remotely:
   ```bash
   git tag -d vX.Y.Z
   git push origin :refs/tags/vX.Y.Z
   ```
2. Delete the release on GitHub (Settings → Releases → Delete)
3. Fix issues and create the tag again

### Binary doesn't work

- Ensure Go version in `.github/workflows/release.yml` matches development version
- Test locally with `make build` before releasing
- Verify target architecture matches your test machine
