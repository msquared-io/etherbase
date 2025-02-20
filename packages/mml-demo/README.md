# MML Etherbase Demo

This project is an example of a server that serves a live
[MML (Metaverse Markup Language)](https://mml.io/) document that uses [React](https://react.dev/).

The MML document is served via a WebSocket and live updated when the code is edited. It can be easily deployed to 
environments that support Node.js and expose ports to the internet.

The react application is inside the [`./mml-document`](./mml-document) folder. 

Edit [`./mml-document/src/index.tsx`](./mml-document/src/index.tsx) whilst the server is running to see your changes reflected live on screen. 

You can find documentation for MML at: 
[`https://mml.io/docs`](https://mml.io/docs).

## Running locally
Making sure you have Node.js installed, run the following from the root of the repository:

```bash
npm install
npm run dev
```

Once the server is running, open `http://localhost:20205` in your browser.
