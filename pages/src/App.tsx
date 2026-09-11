import Markdown from "react-markdown";
import readme from "../README.md?raw";

export const App = () => (
  <main className="min-h-screen bg-background text-foreground">
    <article className="prose prose-neutral dark:prose-invert mx-auto max-w-3xl px-6 py-12">
      <Markdown>{readme}</Markdown>
    </article>
  </main>
);