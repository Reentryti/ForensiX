# ForensiX

> **ForensiX** is a Linux forensic tool designed to **collect, normalize, and analyze large volumes of system artefacts**, with optional assistance from a **LLM** used strictly as an **analysis aid** (correlation, summarization, hypothesis building) — never as a source of truth.

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

* Parsing logs 
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


---

## Use Cases

* User activity timeline reconstruction
* Linux post-incident audit
* Forensic training / CTFs

---

> [!Disclaimer]
> ForensiX is an **analysis assistance tool**.
> It does **not** replace a qualified forensic analyst or legal investigation.

---

## License

This repository is licensed under Apache Licence 2.0 - See [here](LICENSE) for more information.

