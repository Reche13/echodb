import { Button } from "@/components/ui/button";
import Terminal from "@/components/Terminal";
import { ArrowRight, BookOpen, Github } from "lucide-react";
import Link from "next/link";

const HeroSection = () => {
  return (
    <section className="pt-32 pb-20 md:pt-40 md:pb-28">
      <div className="container max-w-5xl mx-auto px-4">
        <div className="grid lg:grid-cols-2 gap-12 lg:gap-16 items-center">
          {/* Left Column - Content */}
          <div className="animate-fade-in" style={{ animationDelay: "0.1s" }}>
            <h1 className="text-4xl md:text-5xl lg:text-6xl font-bold text-foreground tracking-tight leading-tight">
              EchoDB
            </h1>
            <p className="mt-4 text-xl md:text-2xl text-muted-foreground font-medium max-w-[30ch]">
              A lightweight, in-memory database written in Go.
            </p>
            <p className="mt-4 text-muted-foreground leading-relaxed">
              Easy to start. Simple to run. <i>Built for speed</i>.
            </p>

            <div className="mt-8 flex flex-wrap gap-4">
              <Link href="/docs">
                <Button variant="hero" size="lg">
                  Get Started
                  <ArrowRight className="w-4 h-4" />
                </Button>
              </Link>
              <Link
                href="https://github.com/reche13/echodb"
                target="_blank"
                rel="noopener noreferrer"
              >
                <Button variant="hero-ghost" size="lg">
                  <Github className="w-4 h-4" />
                  View Github
                </Button>
              </Link>
            </div>
          </div>

          {/* Right Column - Terminal */}
          <div
            className="animate-fade-in lg:justify-self-end relative"
            style={{ animationDelay: "0.3s" }}
          >
            <Terminal className="w-full" />
          </div>
        </div>
      </div>
    </section>
  );
};

export default HeroSection;
