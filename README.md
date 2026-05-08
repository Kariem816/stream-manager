# Stream Manager

This project aims to be a proxy for clients to connect to WebRTC stream.

## General Idea

WebRTC connection are peer-to-peer (P2P) making it hardware heavy to make calls (or streams) with high number of participants (peers). This architecture also makes every peer IP address visible through the whole network of peers. This project aims to enable secure WebRTC lessons for [Zeta](https://github.com/ZetaLearning/zeta).

## General Workflow

- A teacher will start a lesson broadcast on Zeta.

- Zeta server will create a broadcast instance here.

- Zeta server will make the broadcast url available for students.

- Students will try to connect to the broadcast triggering authorization check.

- If allowed, students will connect to the common broadcast here.

## Consideration

- I guess hosting this program will require setting a subdomain to enable cookie sharing.

- No Database will be needed.

- Stream encoding may be necessary.

- Data channels should be opened or closed by broadcast owner.

- Video storing is not planned yet.