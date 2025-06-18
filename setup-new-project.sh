#!/bin/bash

# Go Gin Boilerplate Setup Script
# This script helps you setup a new project from the boilerplate

set -e

OLD_MODULE="go-gin-boilerplate"
NEW_MODULE="$1"
NEW_APP_NAME="$2"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

print_header() {
    echo -e "${BLUE}================================${NC}"
    echo -e "${BLUE}  Go Gin Boilerplate Setup${NC}"
    echo -e "${BLUE}================================${NC}"
    echo
}

print_success() {
    echo -e "${GREEN}✓${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}⚠${NC} $1"
}

print_error() {
    echo -e "${RED}✗${NC} $1"
}

print_info() {
    echo -e "${BLUE}ℹ${NC} $1"
}

show_usage() {
    echo "Usage: $0 <new-module-name> [app-name]"
    echo
    echo "Examples:"
    echo "  $0 github.com/myorg/my-api"
    echo "  $0 github.com/myorg/my-api \"My API Service\""
    echo
    echo "Arguments:"
    echo "  new-module-name  Go module name (required)"
    echo "  app-name         Application name for docs (optional)"
    exit 1
}

validate_input() {
    if [ -z "$NEW_MODULE" ]; then
        print_error "Module name is required"
        show_usage
    fi

    if [[ ! "$NEW_MODULE" =~ ^[a-zA-Z0-9._/-]+$ ]]; then
        print_error "Invalid module name. Use only letters, numbers, dots, slashes, and hyphens."
        exit 1
    fi

    if [ -z "$NEW_APP_NAME" ]; then
        NEW_APP_NAME=$(basename "$NEW_MODULE")
        print_info "Using '$NEW_APP_NAME' as app name"
    fi
}

backup_files() {
    print_info "Creating backup of original files..."

    if [ ! -d ".backup" ]; then
        mkdir .backup
    fi

    cp go.mod .backup/go.mod.bak 2>/dev/null || true
    cp cmd/main.go .backup/main.go.bak 2>/dev/null || true
    print_success "Backup created in .backup/ directory"
}

update_go_mod() {
    print_info "Updating go.mod file..."

    if [ -f "go.mod" ]; then
        sed -i.bak "s|$OLD_MODULE|$NEW_MODULE|g" go.mod
        rm go.mod.bak 2>/dev/null || true
        print_success "go.mod updated"
    else
        print_warning "go.mod not found"
    fi
}

update_import_paths() {
    print_info "Updating import paths in Go files..."

    # Find all Go files and update import paths
    find . -name "*.go" -type f ! -path "./.backup/*" -exec sed -i.bak "s|$OLD_MODULE|$NEW_MODULE|g" {} +

    # Remove backup files created by sed
    find . -name "*.go.bak" -type f ! -path "./.backup/*" -delete

    print_success "Import paths updated"
}

update_swagger_info() {
    print_info "Updating Swagger documentation info..."

    if [ -f "cmd/main.go" ]; then
        # Update title and description in Swagger comments
        sed -i.bak "s|@title.*|@title           $NEW_APP_NAME API|g" cmd/main.go
        sed -i.bak "s|@description.*|@description     API service for $NEW_APP_NAME|g" cmd/main.go
        rm cmd/main.go.bak 2>/dev/null || true
        print_success "Swagger info updated"
    else
        print_warning "cmd/main.go not found"
    fi
}

update_config_files() {
    print_info "Updating configuration files..."

    # Update database names in config files
    find config/ -name "*.yaml" -type f -exec sed -i.bak "s|chatbot_dev|${NEW_APP_NAME,,}_dev|g" {} + 2>/dev/null || true
    find config/ -name "*.yaml" -type f -exec sed -i.bak "s|chatbot_prod|${NEW_APP_NAME,,}_prod|g" {} + 2>/dev/null || true

    # Remove backup files
    find config/ -name "*.yaml.bak" -type f -delete 2>/dev/null || true

    print_success "Configuration files updated"
}

update_docker_files() {
    print_info "Updating Docker configuration..."

    # Update database names in docker-compose files
    find . -name "docker-compose*.yaml" -type f -exec sed -i.bak "s|chatbot_dev|${NEW_APP_NAME,,}_dev|g" {} + 2>/dev/null || true
    find . -name "docker-compose*.yaml" -type f -exec sed -i.bak "s|chatbot_prod|${NEW_APP_NAME,,}_prod|g" {} + 2>/dev/null || true

    # Remove backup files
    find . -name "docker-compose*.yaml.bak" -type f -delete 2>/dev/null || true

    print_success "Docker files updated"
}

create_dev_config() {
    print_info "Creating development config..."

    if [ ! -d "config/dev" ]; then
        mkdir -p config/dev
    fi

    if [ ! -f "config/dev/config.dev.yaml" ] && [ -f "config/example/config.example.yaml" ]; then
        cp config/example/config.example.yaml config/dev/config.dev.yaml

        # Update the development config
        sed -i.bak "s|env: example|env: development|g" config/dev/config.dev.yaml
        sed -i.bak "s|example_db|${NEW_APP_NAME,,}_dev|g" config/dev/config.dev.yaml
        rm config/dev/config.dev.yaml.bak 2>/dev/null || true

        print_success "Development config created"
    else
        print_warning "Development config already exists or example config not found"
    fi
}

initialize_git() {
    print_info "Checking Git repository..."

    if [ -d ".git" ]; then
        print_warning "Git repository already exists"
        echo "  You may want to:"
        echo "  - git remote set-url origin <your-new-repo-url>"
        echo "  - git add . && git commit -m 'Initial setup from boilerplate'"
    else
        git init
        git add .
        git commit -m "Initial setup from Go Gin boilerplate"
        print_success "Git repository initialized"
        print_info "Don't forget to add your remote origin:"
        echo "  git remote add origin <your-repo-url>"
    fi
}

generate_readme_section() {
    print_info "Updating README with project-specific information..."

    # Create a temporary file with project-specific info
    cat > .project-info.md << EOF

## 🚀 ${NEW_APP_NAME}

This project was created from the Go Gin Boilerplate template.

### Module: \`${NEW_MODULE}\`

### Quick Start for This Project

1. **Setup configuration:**
   \`\`\`bash
   cp config/example/config.example.yaml config/dev/config.dev.yaml
   # Edit config/dev/config.dev.yaml with your settings
   \`\`\`

2. **Start development environment:**
   \`\`\`bash
   make dev
   \`\`\`

3. **Access your API:**
   - API: http://localhost:8080/api/v1
   - Swagger: http://localhost:8080/swagger/index.html
   - Health: http://localhost:8080/api/v1/health

EOF

    print_success "Project-specific README section created in .project-info.md"
}

cleanup() {
    print_info "Cleaning up temporary files..."

    # Remove this setup script
    if [ -f "setup-new-project.sh" ]; then
        rm setup-new-project.sh
        print_success "Setup script removed"
    fi

    # Clean any remaining backup files
    find . -name "*.bak" -type f -delete 2>/dev/null || true
}

print_completion_message() {
    echo
    echo -e "${GREEN}🎉 Project setup completed successfully!${NC}"
    echo
    echo -e "${BLUE}Project Details:${NC}"
    echo -e "  Module: ${GREEN}$NEW_MODULE${NC}"
    echo -e "  App Name: ${GREEN}$NEW_APP_NAME${NC}"
    echo
    echo -e "${BLUE}Next Steps:${NC}"
    echo -e "  1. ${YELLOW}Review and update config/dev/config.dev.yaml${NC}"
    echo -e "  2. ${YELLOW}Update database credentials and names${NC}"
    echo -e "  3. ${YELLOW}Customize Swagger documentation${NC}"
    echo -e "  4. ${YELLOW}Add your remote Git repository${NC}"
    echo -e "  5. ${YELLOW}Start development: make dev${NC}"
    echo
    echo -e "${BLUE}Documentation:${NC}"
    echo -e "  • ${YELLOW}API Docs: http://localhost:8080/swagger/index.html${NC}"
    echo -e "  • ${YELLOW}Health Check: http://localhost:8080/api/v1/health${NC}"
    echo
    echo -e "${GREEN}Happy coding! 🚀${NC}"
}

# Main execution
main() {
    print_header
    validate_input
    backup_files
    update_go_mod
    update_import_paths
    update_swagger_info
    update_config_files
    update_docker_files
    create_dev_config
    generate_readme_section
    initialize_git
    cleanup
    print_completion_message
}

# Check if script is being executed directly
if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
    main "$@"
fi