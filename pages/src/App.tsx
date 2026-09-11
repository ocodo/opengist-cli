import Markdown from "react-markdown";
import readme from "../README.md?raw";

import rehypeHighlight from "rehype-highlight";
import "@highlightjs/cdn-assets/styles/github-dark.min.css";

export const App = () => (
  <main className="min-h-screen bg-background text-foreground">
    <article className="prose prose-neutral dark:prose-invert mx-auto max-w-3xl px-6 py-12 markdown">
      <Markdown rehypePlugins={[rehypeHighlight]}>{readme}</Markdown>
    </article>
  </main>
);