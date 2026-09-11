import Markdown from "react-markdown";
import readme from "../README.md?raw";
import rehypeHighlight from "rehype-highlight";
import "@highlightjs/cdn-assets/styles/github-dark.min.css";
import { CloudDownload } from "lucide-react";

export const App = () => (
  <main className="min-h-screen bg-background/70 text-foreground">
    <article className="prose prose-neutral dark:prose-invert mx-auto max-w-3xl px-6 py-12 markdown">
      <div className="flex flex-row gap-5 items-center justify-center pb-10">
        <div className="text-5xl font-black tracking-tighter m-0 p-0">
          Opengist CLI
        </div>

        <a href="opengist-cli" download="opengist-cli">
          <CloudDownload className="stroke-[1px] w-11 h-11 bg-cyan-800 rounded-full p-1.5 cursor-pointer" />
        </a>
      </div>

      <div className="flex flex-row gap-2 text-sm justify-center font-light">
        Download <a href="opengist-cli" download="opengist-cli" className="decoration-0 hover:decoration-1">opengist-cli</a>
        requires <a href="https://github.com/astral-sh/uv" target="_blank" className="decoration-0 hover:decoration-1">uv,</a> run on Linux or MacOS
      </div>

      <Markdown rehypePlugins={[rehypeHighlight]}>
        {readme.replace("# Opengist-CLI", "")}
      </Markdown>
    </article>
  </main>
);
