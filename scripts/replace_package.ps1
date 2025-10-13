# 替换包名脚本
# 将 runfast-go 替换为 github.com/justatempa/runfast-go

$oldPackage = "runfast-go"
$newPackage = "github.com/justatempa/runfast-go"
$rootDir = Split-Path -Parent $PSScriptRoot

Write-Host "Starting package replacement: $oldPackage -> $newPackage"
Write-Host "Project root directory: $rootDir"

# 获取所有 .go 文件
$goFiles = Get-ChildItem -Path $rootDir -Filter "*.go" -Recurse

$count = 0

foreach ($file in $goFiles) {
    $content = Get-Content -Path $file.FullName -Raw
    
    # 检查文件是否包含旧包名
    if ($content -match [regex]::Escape($oldPackage)) {
        # 替换导入语句中的包名
        $newContent = $content -replace "import\s+\(\s*""$([regex]::Escape($oldPackage))", "import (`"$newPackage"
        $newContent = $newContent -replace "import\s+""$([regex]::Escape($oldPackage))", "import `"$newPackage"
        
        # 替换导入语句中的子包
        $newContent = $newContent -replace """$([regex]::Escape($oldPackage))/", "`"$newPackage/"
        
        # 写回文件
        Set-Content -Path $file.FullName -Value $newContent
        
        Write-Host "Updated: $($file.FullName)"
        $count++
    }
}

# 更新 go.mod 文件
$goModPath = Join-Path -Path $rootDir -ChildPath "go.mod"
if (Test-Path $goModPath) {
    $goModContent = Get-Content -Path $goModPath -Raw
    $newGoModContent = $goModContent -replace "module\s+$([regex]::Escape($oldPackage))", "module $newPackage"
    
    if ($goModContent -ne $newGoModContent) {
        Set-Content -Path $goModPath -Value $newGoModContent
        Write-Host "Updated go.mod file"
        $count++
    }
}

Write-Host "Replacement completed, updated $count files"