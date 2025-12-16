# AnaLinux

> **ForensicLM** is a Linux forensic tool designed to **collect, normalize, and analyze large volumes of system artefacts**, with optional assistance from a **LLM** used strictly as an **analysis aid** (correlation, summarization, hypothesis building) — never as a source of truth.

---

## Project Goals

* Facilitate analysis of **large-scale Linux forensic data**
* Enforce a **forensic-safe approach** (reproducible, traceable, non-intrusive)
* Strictly separate:

  * **deterministic evidence collection**
  * **LLM-assisted interpretation**
* Provide a solid foundation for:

  * an **open-source forensic tool**
  * **forensic CTFs** or incident response

---

## High-Level Architecture

```
[ Disk image / RO-mounted system ]
                ↓
          Go collectors
                ↓
      Normalized events
           (JSONL)
                ↓
     Deterministic correlation
                ↓
        LLM assistance
                ↓
     Reports / hypotheses
```

---

## Features (MVP)

### Collection

* Parsing `/var/log/auth.log`
* Extraction of SSH attempts 
* JSON Lines output 

### Normalization

* Unified event schema
* ISO 8601 timestamps
* Precise evidence references 

### Analysis

* Time-based filtering
* Grouping by user / IP
* Event batching for LLM consumption

---

## Role of the LLM

The LLM is used strictly to:

* correlate already-extracted events
* detect suspicious patterns
* propose **evidence-based hypotheses**
* generate human-readable summaries

### Strict Constraints

* no fact fabrication
* no data modification
* every claim must reference `event_id`s

---

## Technology 

* **Language**: Go
* **Formats**: JSON, JSON Lines
* **Hashing**: SHA-256
* **LLM**: local

---

## Quick Start 

```bash
go build ./cmd/forensiclm
./forensiclm collect --root /mnt/image
```

---

## Use Cases

* SSH intrusion analysis
* User activity timeline reconstruction
* Linux post-incident audit
* Forensic training / CTFs

---

## ⚠️ Disclaimer

ForensicLM is an **analysis assistance tool**.
It does **not** replace a qualified forensic analyst or legal investigation.

---

## License

MIT License

