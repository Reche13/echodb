import { cn } from "@/lib/utils";

interface TerminalProps {
  className?: string;
}

const Terminal = ({ className }: TerminalProps) => {
  return (
    <div
      className={cn(
        "rounded-xl overflow-hidden shadow-card border border-border/50",
        className
      )}
    >
      {/* Terminal Header */}
      <div className="bg-code-bg px-4 py-3 flex items-center gap-2">
        <div className="flex gap-1.5">
          <div className="w-3 h-3 rounded-full bg-red-400/80" />
          <div className="w-3 h-3 rounded-full bg-yellow-400/80" />
          <div className="w-3 h-3 rounded-full bg-green-400/80" />
        </div>
        <span className="text-code-text/60 text-xs font-mono ml-2">echodb</span>
      </div>

      {/* Terminal Body */}
      <div className="bg-code-bg px-4 py-4 font-mono text-sm leading-relaxed">
        <pre className="text-terminal-orange">
          {`███████╗ ██████╗██╗  ██╗ ██████╗ 
██╔════╝██╔════╝██║  ██║██╔═══██╗
█████╗  ██║     ███████║██║   ██║
██╔══╝  ██║     ██╔══██║██║   ██║
███████╗╚██████╗██║  ██║╚██████╔╝
╚══════╝ ╚═════╝╚═╝  ╚═╝ ╚═════╝`}
        </pre>
        <div className="mt-3 text-code-text">
          <span className="text-cyan-400">EchoDB</span>
          <span className="text-code-text/60"> — In-Memory Data Store</span>
        </div>
        <div className="text-code-text/60 mt-1">
          Version: <span className="text-green-500">0.1.0</span>
        </div>
        <div className="text-code-text/60 mt-2">
          Listening on{" "}
          <span className="text-terminal-yellow">0.0.0.0:6380</span>
        </div>
        <div className="mt-3 flex items-center">
          <span className="text-terminal-green">❯</span>
          <span className="w-2 h-4 bg-code-text/80 ml-1 animate-terminal-blink" />
        </div>
      </div>
    </div>
  );
};

export default Terminal;
