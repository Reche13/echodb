import CodeBlock from "@/components/CodeBlock";

const QuickStartSection = () => {
  return (
    <section id="quickstart" className="py-20 bg-surface/50 scroll-mt-20">
      <div className="container max-w-6xl mx-auto px-4">
        <div className="text-center mb-12">
          <h2 className="text-3xl md:text-4xl font-bold text-foreground">
            Start EchoDB in Seconds
          </h2>
          <p className="mt-4 text-muted-foreground">
            No setup. No config. Just run.
          </p>
        </div>

        <div className="grid md:grid-cols-2 gap-6 max-w-4xl mx-auto">
          <div className="animate-fade-in" style={{ animationDelay: "0.1s" }}>
            <p className="text-sm font-medium text-muted-foreground mb-3">
              Pull & Run with Docker
            </p>
            <CodeBlock
              title="terminal"
              code={`docker pull rechesoares/echodb:0.1.0
docker run -p 6380:6380 rechesoares/echodb:0.1.0`}
            />
          </div>

          <div className="animate-fade-in" style={{ animationDelay: "0.2s" }}>
            <p className="text-sm font-medium text-muted-foreground mb-3">
              Verify It Works
            </p>
            <CodeBlock
              title="terminal"
              code={`redis-cli -p 6380
PING
PONG`}
            />
          </div>
        </div>
      </div>
    </section>
  );
};

export default QuickStartSection;
