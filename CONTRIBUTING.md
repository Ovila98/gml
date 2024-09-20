# Contributing to GML

We welcome contributions to the GML project! Whether you want to report issues, suggest new features, or submit pull requests, we appreciate all kinds of help from the community.

## How to Contribute

### 1. Reporting Issues

If you encounter any bugs, issues, or have feature requests, please open an issue on GitHub:

1. Go to the **Issues** tab.
2. Click on **New Issue**.
3. Provide a clear and concise description of the problem or feature request.

Please include:

- A clear title.
- Steps to reproduce the problem (if applicable).
- Expected behavior vs. actual behavior.
- The environment where the issue occurs (Go version, OS, etc.).
- Possible solutions or suggestions.

### 2. Submitting Code Contributions (Pull Requests)

#### Step 1: Fork the Repository

To make changes, fork the repository on GitHub and clone it to your local machine:

```bash
git clone https://github.com/your-username/gml.git
```

#### Step 2: Create a Branch

Create a new branch for your changes:

```bash
git checkout -b feature-name
```

Name your branch after the feature you are adding or the bug you are fixing.

#### Step 3: Code

Make your changes or additions to the code. We expect the following:

- Follow the Go [code formatting and style guidelines](https://golang.org/doc/effective_go.html).
- Ensure your code is well-documented with comments.
- Write tests for any new functionality or changes. We use Go’s built-in `testing` package.

#### Step 4: Commit Changes

Commit your changes with a descriptive message:

```bash
git add .
git commit -m "Add feature X"
```

#### Step 5: Run Tests

Before submitting, ensure all tests pass:

```bash
go test ./...
```

#### Step 6: Submit a Pull Request

1. Push your branch to your GitHub fork:
   ```bash
   git push origin feature-name
   ```
2. Go to the original repository on GitHub.
3. Click on **New Pull Request**.
4. Provide a clear description of your changes and link to any relevant issues.
5. Submit the pull request.

### 3. Code Review

All pull requests will be reviewed by the maintainers. During this process, you may receive feedback or requests for changes. We expect:

- Constructive communication.
- Willingness to make requested updates or changes.
- Patience as maintainers review your code.

Once the changes are approved, your code will be merged into the main branch.

### 4. Coding Standards

Please ensure that your contributions adhere to the following:

- Follow Go’s standard library style and practices, as outlined in [Effective Go](https://golang.org/doc/effective_go.html).
- Use clear and meaningful variable and function names.
- Keep functions focused and small, avoiding unnecessary complexity.
- Comment exported functions, types, and variables.
- Format code using Go’s `go fmt` tool.

### 5. Writing Tests

We value robust test coverage to ensure the stability of the `gml` project. For all new features or bug fixes, make sure to:

- Write tests in the `_test.go` files.
- Use `t.Run` to structure sub-tests.
- Aim for clear, readable test cases that cover edge cases.
- Run the tests locally before submitting.

### 6. Feature Requests

If you have ideas for new features, feel free to:

1. Open a GitHub issue outlining your feature suggestion.
2. Provide a clear and detailed description of the proposed functionality.
3. Engage in discussions with maintainers and contributors.

### Code of Conduct

We expect all contributors to adhere to our [Code of Conduct](CODE_OF_CONDUCT.md), fostering a positive and constructive environment for everyone involved in the project.

---

Thank you for contributing to **GML**! Your efforts help make this project better for everyone.
