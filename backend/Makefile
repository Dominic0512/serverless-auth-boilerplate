# Define the command for generating documentation
GEN_DOCS = swag init

# Define the output directories and source directories for each service
AUTH_OUTPUT = cmd/auth/docs
AUTH_DIRS = ./cmd/auth/gin,./controller/auth

USER_OUTPUT = cmd/user/docs
USER_DIRS = ./cmd/user/gin,./controller/user

# Target to generate documentation for auth and user services
gen-docs: auth-docs user-docs

# Rule to generate documentation for auth service
auth-docs:
	$(GEN_DOCS) -o $(AUTH_OUTPUT) -d $(AUTH_DIRS)

# Rule to generate documentation for user service
user-docs:
	$(GEN_DOCS) -o $(USER_OUTPUT) -d $(USER_DIRS)

# Clean up generated documentation (optional)
clean:
	rm -rf $(AUTH_OUTPUT)/* $(USER_OUTPUT)/*