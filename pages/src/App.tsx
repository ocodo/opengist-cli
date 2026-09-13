import Markdown from "react-markdown";
import readme from "../README.md?raw";
import rehypeHighlight from "rehype-highlight";
import "@highlightjs/cdn-assets/styles/github-dark.min.css";
import { Terminal } from "lucide-react";
import { GithubIcon } from "@/components/GHIcon";
import { FaApple, FaLinux, FaWindows } from "react-icons/fa";

const RELEASE_URL = "https://github.com/ocodo/opengist-cli/releases/download/1.0.0"

const PLATFORMS = [
  { label: "Linux (amd64)", file: "opengist-cli-linux-amd64", icon: <FaLinux /> },
  { label: "Linux (arm64)", file: "opengist-cli-linux-arm64", icon: <FaLinux /> },
  { label: "macOS (Apple Silicon)", file: "opengist-cli-darwin-arm64", icon: <FaApple /> },
  { label: "macOS (Intel)", file: "opengist-cli-darwin-amd64", icon: <FaApple /> },
  { label: "Windows (x64)", file: "opengist-cli-windows-amd64.exe", icon: <FaWindows /> },
];

export const App = () => (
  <main className="min-h-screen bg-background/70 text-foreground">
    <article className="prose  dark:prose-invert mx-auto max-w-3xl px-6 py-12 markdown">
      <a href="https://github.com/ocodo/opengist-cli" target="_blank" rel="noreferrer">
        <svg xmlns="http://www.w3.org/2000/svg" className="fixed top-0 right-0 w-30 h-30" viewBox="0 0 100 100">
          <polygon points="0,0 100,0 100,100" fill="#00000040" />
        </svg>
        <GithubIcon className="m-2.5 fixed top-0 right-0" />
      </a>

      <div className="flex flex-row gap-4 justify-center items-center text-5xl font-black tracking-tighter pb-6">
        <Terminal className="h-16 w-16" />
        <div>Opengist CLI</div>
      </div>

      <div className="not-prose flex flex-col items-center gap-4">
	<div className="grid grid-cols-2 gap-5">
	  {PLATFORMS.map((p) => (
	    <a
	      key={p.file}
	      href={`${RELEASE_URL}/downloads/${p.file}`}
	      download={p.file}
	      className="rounded-xl border border-foreground/10 flex flex-col items-center px-3 py-1.5 hover:bg-background/25 transition-colors last:col-span-2 last:justify-self-center last:w-1/2"
	    >
	      {p.icon}
	      {p.label}
	    </a>
	  ))}
	</div>
      </div>

      <Markdown rehypePlugins={[rehypeHighlight]}>
        {readme}
      </Markdown>
    </article>
  </main>
);
