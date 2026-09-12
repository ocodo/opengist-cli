import Markdown from "react-markdown";
import readme from "../README.md?raw";
import rehypeHighlight from "rehype-highlight";
import "@highlightjs/cdn-assets/styles/github-dark.min.css";
import { CloudDownload } from "lucide-react";
import { GithubIcon } from "@/components/GHIcon";

export const App = () => (
  <main className="min-h-screen bg-background/70 text-foreground">
    <article className="prose prose-neutral dark:prose-invert mx-auto max-w-3xl px-6 py-12 markdown">

      <a href="https://github.com/ocodo/opengist-cli" target='_blank'>
        <svg xmlns="http://www.w3.org/2000/svg" className="fixed top-0 right-0 w-30 h-30" viewBox="0 0 100 100">
          <polygon points="0,0 100,0 100,100" fill="#00000040" />
        </svg>
        <GithubIcon className="m-2.5 fixed top-0 right-0" />
      </a>
      <a href="opengist-cli" download="opengist-cli">
        <CloudDownload className="fixed top-4 left-4 stroke-[1px] w-8 h-8 bg-foreground text-background rounded-full p-1.5 cursor-pointer" />
      </a>

      <div className="flex flex-col items-center text-5xl font-black tracking-tighter py-6">
        <div>
          Opengist CLI
        </div>
      </div>

      <div className="flex flex-wrap gap-2 text-sm justify-center font-light">
        Download <a href="opengist-cli" download="opengist-cli" className="decoration-0 hover:decoration-1">opengist-cli</a>
        requires <a href="https://github.com/astral-sh/uv" target="_blank" className="decoration-0 hover:decoration-1">uv,</a> runs on all platforms uv does
      </div>

      <Markdown rehypePlugins={[rehypeHighlight]}>
        {readme.replace("# Opengist-CLI", "")}
      </Markdown>
    </article>
  </main>
);
