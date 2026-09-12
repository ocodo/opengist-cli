import React from "react";
import ReactDOM from "react-dom/client";
import fontUrl from "./assets/OcodoMonoDotZeroNerdFont-Light.woff2?url";
import { App } from "./App";
import "./index.css";

// Inject font-face with guaranteed asset path
const style = document.createElement("style");
style.textContent = `
  @font-face {
    src: url("${fontUrl}");
    font-family: "OMDZero";
    font-weight: 300;
    font-style: normal;
  }
`;
document.head.appendChild(style);

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
);
