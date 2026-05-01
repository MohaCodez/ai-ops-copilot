# AI Ops Copilot

Day 1:
- Go API
- Health endpoint
- Basic structure

Day 2
- Swapped default mux for Chi router
- Added global logger middleware (method + path + latency)
- Added Recoverer middleware (panics → 500, server stays alive)
- Built POST /ask with typed request/response structs
- Input validation (bad JSON → 400, empty query → 400)
- Stubbed answer response (RAG hookup comes later)
- Clean folder structure: handlers/, middleware/, cmd/