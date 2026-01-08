import { LucideIcon } from "lucide-react";
import { cn } from "@/lib/utils";
import CodeBlock from "./CodeBlock";

interface InstallCardProps {
  icon: LucideIcon;
  title: string;
  description: string;
  code: string;
  recommended?: boolean;
  className?: string;
}

const InstallCard = ({
  icon: Icon,
  title,
  description,
  code,
  recommended,
  className,
}: InstallCardProps) => {
  return (
    <div
      className={cn(
        "rounded-xl border border-border bg-card overflow-hidden",
        "transition-all duration-300 hover:shadow-card hover:border-primary/20",
        className
      )}
    >
      <div className="p-6 pb-4">
        <div className="flex items-start justify-between mb-3">
          <div className="flex items-center gap-3">
            <div className="w-10 h-10 rounded-lg bg-primary/10 flex items-center justify-center">
              <Icon className="w-5 h-5 text-primary" />
            </div>
            <div>
              <h3 className="font-semibold text-foreground">{title}</h3>
              {recommended && (
                <span className="text-xs text-primary font-medium">
                  Recommended
                </span>
              )}
            </div>
          </div>
        </div>
        <p className="text-muted-foreground text-sm">{description}</p>
      </div>
      <div className="px-4 pb-4">
        <CodeBlock code={code} className="shadow-none border-0" />
      </div>
    </div>
  );
};

export default InstallCard;
