import Markdown from "react-markdown";
import readme from "../README.md?raw";
import rehypeHighlight from "rehype-highlight";
import "@highlightjs/cdn-assets/styles/github-dark.min.css";

export const App = () => (
  <main className="min-h-screen bg-background/70 text-foreground">
    <article className="prose prose-neutral dark:prose-invert mx-auto max-w-3xl px-6 py-12 markdown">
      <h1 className="text-5xl font-black tracking-tighter">Opengist CLI</h1>
      <Markdown rehypePlugins={[rehypeHighlight]}>
        {readme.replace("# Opengist-CLI", "")}
      </Markdown>
    </article>
  </main>
);