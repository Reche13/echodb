import InstallCard from "@/components/InstallCard";
import { Container, Code2 } from "lucide-react";

const InstallSection = () => {
  return (
    <section id="install" className="py-20 scroll-mt-20">
      <div className="container max-w-6xl mx-auto px-4">
        <div className="text-center mb-12">
          <h2 className="text-3xl md:text-4xl font-bold text-foreground">
            Install It Your Way
          </h2>
          <p className="mt-4 text-muted-foreground">
            Choose the installation method that works best for you.
          </p>
        </div>

        <div className="grid md:grid-cols-2 gap-6 max-w-4xl mx-auto">
          <div className="animate-fade-in" style={{ animationDelay: "0.1s" }}>
            <InstallCard
              icon={Container}
              title="Docker"
              description="The easiest way to run EchoDB."
              code="docker run -p 6380:6380 rechesoares/echodb:0.1.0"
              recommended
            />
          </div>

          <div className="animate-fade-in" style={{ animationDelay: "0.2s" }}>
            <InstallCard
              icon={Code2}
              title="Build From Source"
              description="For contributors and tinkerers."
              code={`git clone https://github.com/reche13/echodb.git
cd echodb
go build cmd/main.go`}
            />
          </div>
        </div>
      </div>
    </section>
  );
};

export default InstallSection;
