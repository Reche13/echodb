import FeatureCard from "@/components/FeatureCard";
import { Zap, Radio, Code2, Wrench } from "lucide-react";

const features = [
  {
    icon: Zap,
    title: "In-Memory Speed",
    description: "Built for fast reads and writes with minimal overhead.",
  },
  {
    icon: Radio,
    title: "RESP Compatible",
    description: "Works with existing Redis clients.",
  },
  {
    icon: Code2,
    title: "Written in Go",
    description: "Simple, predictable performance.",
  },
];

const FeaturesSection = () => {
  return (
    <section id="features" className="py-20 scroll-mt-20">
      <div className="container max-w-5xl mx-auto px-4">
        <div className="text-center mb-12">
          <h2 className="text-3xl md:text-4xl font-bold text-foreground">
            Why EchoDB?
          </h2>
          <p className="mt-4 text-muted-foreground max-w-2xl mx-auto">
            A minimal, focused approach to in-memory data storage.
          </p>
        </div>

        <div className="grid sm:grid-cols-1 lg:grid-cols-3 gap-6">
          {features.map((feature, index) => (
            <FeatureCard
              key={feature.title}
              icon={feature.icon}
              title={feature.title}
              description={feature.description}
              className="animate-fade-in"
              style={
                {
                  animationDelay: `${0.1 + index * 0.1}s`,
                } as React.CSSProperties
              }
            />
          ))}
        </div>
      </div>
    </section>
  );
};

export default FeaturesSection;
