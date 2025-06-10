# AGENTS.md

This document defines the set of agents your Codex-powered system will use to build a cloud gaming service capable of streaming and playing PS2 games. Each agent has a clear responsibility and communicates with others via well-defined interfaces.

---

## Table of Contents

1. [Overview](#overview)  
2. [Agent Definitions](#agent-definitions)  
   - [1. AuthenticationAgent](#1-authenticationagent)  
   - [2. GameSessionManager](#2-gamesessionmanager)  
   - [3. InputProcessorAgent](#3-inputprocessoragent)  
   - [4. VideoEncoderAgent](#4-videoencoderagent)  
   - [5. NetworkTransportAgent](#5-networktransportagent)  
   - [6. ResourceAllocatorAgent](#6-resourceallocatoragent)  
   - [7. MonitoringAgent](#7-monitoringagent)  
   - [8. UIOrchestratorAgent](#8-uiorchestratoragent)  
3. [Inter-Agent Protocol](#inter-agent-protocol)  
4. [Deployment Diagram](#deployment-diagram)  
5. [Next Steps](#next-steps)  

---

## Overview

Your cloud gaming service consists of specialized agents that collaborate to:

- Authenticate users and sessions  
- Launch and manage PS2 emulator instances  
- Process controller/keyboard inputs  
- Encode and stream video/audio to clients  
- Allocate compute and network resources  
- Monitor health, performance, and usage  
- Orchestrate front-end UI updates  

Use Codex to implement each agent in modular code, with clear interfaces and unit tests.

---

## Agent Definitions

### 1. AuthenticationAgent

- **Responsibility**  
  - Validate user credentials (JWT, OAuth)  
  - Issue session tokens  
  - Refresh tokens on expiration  

- **Inputs**  
  - `username`, `password` or `refresh_token`  

- **Outputs**  
  - `access_token`, `refresh_token`, `expires_in`  

- **APIs / Tools**  
  - Identity Provider (e.g. Auth0, Cognito) SDK  
  - In-memory or Redis cache for token blacklisting  

- **Sample Prompt for Codex**  
  > “Generate a Node.js module `AuthenticationAgent` that exposes `login()`, `refreshToken()`, and `verifyToken()` using Auth0’s SDK.”

---

### 2. GameSessionManager

- **Responsibility**  
  - Spawn and manage PS2 emulator processes (e.g., PCSX2 in headless mode)  
  - Track session lifecycle: start, pause, stop, teardown  
  - Persist session metadata  

- **Inputs**  
  - `access_token`, `game_id`, `session_config`  

- **Outputs**  
  - `session_id`, `emulator_endpoint`, `status`  

- **APIs / Tools**  
  - Docker API or Kubernetes client to schedule emulator containers  
  - Database (PostgreSQL) for session records  

- **Sample Prompt for Codex**  
  > “Write a Python `GameSessionManager` class that uses the Kubernetes Python client to launch a PS2 emulator pod, and records the session in PostgreSQL.”

---

### 3. InputProcessorAgent

- **Responsibility**  
  - Receive raw controller/keyboard events from clients  
  - Normalize and inject them into the emulator process  

- **Inputs**  
  - WebSocket messages carrying input frames  

- **Outputs**  
  - OS-level input events (e.g., via `uinput` on Linux)  

- **APIs / Tools**  
  - `ws` or Socket.IO for WebSocket handling  
  - Linux `uinput` library  

- **Sample Prompt for Codex**  
  > “Implement `InputProcessorAgent` in Go: open a WebSocket, parse JSON input frames, and emit uinput events to the emulator’s virtual device.”

---

### 4. VideoEncoderAgent

- **Responsibility**  
  - Capture emulator’s framebuffer output  
  - Encode video (H.264) and audio (AAC) in real time  
  - Package into fragmented MP4 or low-latency HLS  

- **Inputs**  
  - Raw video + audio streams from the emulator container  

- **Outputs**  
  - RTP/HLS segments or WebRTC media streams  

- **APIs / Tools**  
  - FFmpeg CLI or libavcodec bindings  
  - WebRTC libraries (e.g., Pion for Go, aiortc for Python)  

- **Sample Prompt for Codex**  
  > “Generate a Dockerized `VideoEncoderAgent` that runs FFmpeg in realtime mode, reads from `/dev/shm/framebuffer`, and serves a WebRTC endpoint using aiortc.”

---

### 5. NetworkTransportAgent

- **Responsibility**  
  - Handle low-latency transport of media and input  
  - Manage NAT traversal, TURN/STUN if needed  
  - Route packets between client and encoder/input agents  

- **Inputs**  
  - Encoded media packets, input acknowledgment packets  

- **Outputs**  
  - Forwarded packets over UDP/TCP/WebRTC channels  

- **APIs / Tools**  
  - `coturn` server configuration  
  - WebRTC datachannels for input fallback  

- **Sample Prompt for Codex**  
  > “Create a TypeScript `NetworkTransportAgent` that integrates with `simple-peer` to negotiate WebRTC between browser clients and the media server.”

---

### 6. ResourceAllocatorAgent

- **Responsibility**  
  - Monitor cluster load (CPU, GPU, RAM)  
  - Scale emulator nodes up/down based on demand  
  - Enforce per-user quotas  

- **Inputs**  
  - Metrics from Prometheus or cloud provider APIs  

- **Outputs**  
  - Kubernetes HorizontalPodAutoscaler adjustments  
  - Alerts when capacity is exhausted  

- **APIs / Tools**  
  - Prometheus HTTP API  
  - Kubernetes HPA controller  

- **Sample Prompt for Codex**  
  > “Write a Python `ResourceAllocatorAgent` that queries Prometheus every 30 seconds and updates a Kubernetes HPA resource accordingly.”

---

### 7. MonitoringAgent

- **Responsibility**  
  - Collect logs, performance metrics, session analytics  
  - Trigger alerts on errors, high latency, or crashes  
  - Expose dashboards (Grafana)  

- **Inputs**  
  - Log files, Prometheus metrics, session events  

- **Outputs**  
  - Alerts via Slack/Email/PagerDuty  
  - Metrics endpoint for dashboards  

- **APIs / Tools**  
  - Grafana provisioning API  
  - Alertmanager integration  

- **Sample Prompt for Codex**  
  > “Implement a Go `MonitoringAgent` that scrapes log files for error patterns and sends Slack notifications when detected.”

---

### 8. UIOrchestratorAgent

- **Responsibility**  
  - Serve the web client or Electron app  
  - Coordinate signaling for matchmaking, session join/leave  
  - Render game lobby, player controls, and status  

- **Inputs**  
  - HTTP requests from browser/Electron frontend  
  - WebSocket messages for real-time updates  

- **Outputs**  
  - HTML/JS/CSS bundles, REST & WebSocket responses  

- **APIs / Tools**  
  - React/Next.js for front end  
  - Express/Koa or FastAPI for backend  

- **Sample Prompt for Codex**  
  > “Generate a Next.js `UIOrchestratorAgent` with API routes `/api/login`, `/api/start-session`, and a React component for the game lobby.”

---

## Inter-Agent Protocol

All agents communicate over JSON-RPC or REST/WebSocket with the following conventions:

- **Authentication**: Bearer tokens in the `Authorization` header.  
- **Session Context**: Every request includes `session_id` and `user_id`.  
- **Error Handling**: Use standardized error codes `{ code: number, message: string }`.  
- **Logging**: Agents emit structured logs in JSON to stdout.  

---

## Deployment Diagram

[Client Browser]
│ WebRTC / WebSocket
↓
[NetworkTransportAgent] ←→ [VideoEncoderAgent] ←→ [GameSessionManager] ←→ [InputProcessorAgent]
↑ │
│ ↓
[UIOrchestratorAgent] ←→ [AuthenticationAgent] [ResourceAllocatorAgent]
│
↓
[MonitoringAgent]

---

## Next Steps

1. Scaffold each agent folder with language‐specific boilerplate.  
2. Write interface definitions (TypeScript `.d.ts` or Python interfaces).  
3. Create Codex prompts based on the “Sample Prompt” for each agent.  
4. Build, test, and iterate—start with AuthenticationAgent and GameSessionManager.  
5. Integrate end-to-end streaming from emulator to client.