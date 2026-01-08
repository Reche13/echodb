import { LucideIcon } from "lucide-react";
import { cn } from "@/lib/utils";

interface FeatureCardProps {
  icon: LucideIcon;
  title: string;
  description: string;
  className?: string;
  style?: React.CSSProperties;
}

const FeatureCard = ({
  icon: Icon,
  title,
  description,
  className,
  style,
}: FeatureCardProps) => {
  return (
    <div
      className={cn(
        "group p-6 rounded-xl border border-border bg-card",
        "transition-all duration-300",
        "hover:shadow-card hover:border-primary/20 hover:-translate-y-0.5",
        className
      )}
      style={style}
    >
      <div className="w-10 h-10 rounded-lg bg-primary/10 flex items-center justify-center mb-4 group-hover:bg-primary/15 transition-colors">
        <Icon className="w-5 h-5 text-primary" />
      </div>
      <h3 className="text-lg font-semibold text-foreground mb-2">{title}</h3>
      <p className="text-muted-foreground text-sm leading-relaxed">
        {description}
      </p>
    </div>
  );
};

export default FeatureCard;
