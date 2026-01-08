import { Button } from "@/components/ui/button";
import { BookOpen, Github } from "lucide-react";

const CTASection = () => {
  return (
    <section className="py-20 bg-surface/50">
      <div className="container max-w-6xl mx-auto px-4">
        <div className="text-center max-w-2xl mx-auto">
          <h2 className="text-3xl md:text-4xl font-bold text-foreground">
            Explore the Docs
          </h2>
          <p className="mt-4 text-muted-foreground leading-relaxed">
            Learn how EchoDB works under the hood and how to use it effectively.
          </p>

          <div className="mt-8 flex flex-wrap justify-center gap-4">
            <a href="/docs" rel="noopener noreferrer">
              <Button variant="hero" size="lg">
                <BookOpen className="w-4 h-4" />
                Read the Documentation
              </Button>
            </a>
            <a
              href="https://github.com/reche13/echodb"
              target="_blank"
              rel="noopener noreferrer"
            >
              <Button variant="hero-ghost" size="lg">
                <Github className="w-4 h-4" />
                View on GitHub
              </Button>
            </a>
          </div>
        </div>
      </div>
    </section>
  );
};

export default CTASection;
