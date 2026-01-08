"use client";
import { Copy, Check } from "lucide-react";
import { useState } from "react";
import { cn } from "@/lib/utils";

interface CodeBlockProps {
  title?: string;
  code: string;
  language?: string;
  className?: string;
}

const CodeBlock = ({ title, code, language, className }: CodeBlockProps) => {
  const [copied, setCopied] = useState(false);

  const handleCopy = async () => {
    await navigator.clipboard.writeText(code);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  return (
    <div
      className={cn(
        "rounded-xl overflow-hidden shadow-card border border-border/50",
        className
      )}
    >
      {/* Header */}
      <div className="bg-code-bg px-4 py-2.5 flex items-center justify-between border-b border-white/5">
        <div className="flex items-center gap-2">
          <div className="flex gap-1.5">
            <div className="w-2.5 h-2.5 rounded-full bg-red-400/60" />
            <div className="w-2.5 h-2.5 rounded-full bg-yellow-400/60" />
            <div className="w-2.5 h-2.5 rounded-full bg-green-400/60" />
          </div>
          {title && (
            <span className="text-code-text/50 text-xs font-mono ml-2">
              {title}
            </span>
          )}
        </div>
        <button
          onClick={handleCopy}
          className="text-code-text/40 hover:text-code-text/80 transition-colors p-1"
          aria-label="Copy code"
        >
          {copied ? (
            <Check className="w-4 h-4" />
          ) : (
            <Copy className="w-4 h-4" />
          )}
        </button>
      </div>

      {/* Code Body */}
      <div className="bg-code-bg px-4 py-4 overflow-x-auto">
        <pre className="font-mono text-sm text-code-text leading-relaxed whitespace-pre">
          {code}
        </pre>
      </div>
    </div>
  );
};

export default CodeBlock;
