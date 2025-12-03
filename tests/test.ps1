Write-Host "=== Blue Carbon MRV Blockchain Project ===" -ForegroundColor Green
Write-Host ""

# Check prerequisites
Write-Host "Checking prerequisites..." -ForegroundColor Yellow

# Check Go
try {
    $goVersion = go version 2>$null
    Write-Host "✓ Go: $goVersion" -ForegroundColor Green
} catch {
    Write-Host "✗ Go not installed" -ForegroundColor Red
}

# Check Node.js
try {
    $nodeVersion = node --version 2>$null
    Write-Host "✓ Node.js: $nodeVersion" -ForegroundColor Green
} catch {
    Write-Host "✗ Node.js not installed" -ForegroundColor Red
}

# Check Docker
try {
    $dockerVersion = docker --version 2>$null
    Write-Host "✓ Docker: $dockerVersion" -ForegroundColor Green
} catch {
    Write-Host "✗ Docker not installed" -ForegroundColor Red
}

# Check Git
try {
    $gitVersion = git --version 2>$null
    Write-Host "✓ Git: $gitVersion" -ForegroundColor Green
} catch {
    Write-Host "✗ Git not installed" -ForegroundColor Red
}

Write-Host ""
Write-Host "Project structure created successfully!" -ForegroundColor Green
Write-Host "Next steps:" -ForegroundColor Cyan
Write-Host "1. Open VS Code in this directory"
Write-Host "2. Run: cd api && node server.js"
Write-Host "3. Visit: http://localhost:3000/health"