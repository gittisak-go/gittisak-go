# Custom GitHub Copilot Agents

This directory contains custom agent configurations for GitHub Copilot that provide specialized expertise and guidance.

## Available Agents

### 1. MIT Software Architecture Professor (`mit-professor.md`)
An emeritus MIT professor specializing in software engineering and code architecture.

**Expertise Areas:**
- Clean Code & Architecture (SOLID, Design Patterns)
- Python (FastAPI, Django, async programming)
- React & React Native
- API Design (REST, GraphQL)
- DevOps (Docker, CI/CD, GitHub Actions)
- Cloud Platforms (AWS, Azure, GCP)
- AI API Integration (Claude, OpenAI)

**Best For:**
- Code generation with industry best practices
- Architectural decisions and trade-off analysis
- Git workflows and CI/CD setup
- Performance optimization
- Security reviews
- Teaching complex concepts with the "why" behind decisions

**Response Style:**
- Precise and authoritative
- Complete working code examples
- Proactive (anticipates security and performance issues)
- Focuses on practical implementation
- Explains reasoning, not just solutions

### 2. MIT Software Architecture Professor (Thai) (`mit-professor-th.md`)
Same as above but with Thai language documentation.

## How to Use

These custom agents are automatically available in GitHub Copilot when you work in this repository. To invoke them:

1. **In GitHub Copilot Chat:**
   - Use `@mit-professor` (or reference the agent by name)
   - Ask your question or request code assistance
   - The agent will respond according to its specialized persona and expertise

2. **For Code Reviews:**
   - The agent can provide in-depth architectural reviews
   - Security and performance analysis
   - Best practice recommendations

3. **For Learning:**
   - Ask for explanations of complex concepts
   - Request breakdown of architectural decisions
   - Learn about trade-offs in different approaches

## Example Prompts

```
@mit-professor Can you help me design a FastAPI application with authentication?

@mit-professor Review this code for security vulnerabilities and performance issues

@mit-professor Explain the trade-offs between REST and GraphQL for my use case

@mit-professor Show me how to set up a CI/CD pipeline with GitHub Actions for a Python project

@mit-professor What's the best caching strategy for a read-heavy API?
```

## Agent Configuration Format

Each agent is defined in a Markdown file with:
- **Persona**: Who the agent is and their background
- **Mission**: Core objectives and goals
- **Technical Expertise**: Specific technologies and domains
- **Response Style**: How the agent communicates
- **Best Practices**: Principles and guidelines the agent follows

## Contributing

To add a new custom agent:
1. Create a new `.md` file in this directory
2. Define the persona, expertise, and response style
3. Test the agent with various prompts
4. Update this README with the new agent details

## Notes

- Agents have specialized knowledge but can handle general software engineering questions
- Responses focus on production-ready, maintainable code
- Security and performance are considered proactively
- Explanations emphasize understanding "why" not just "how"
