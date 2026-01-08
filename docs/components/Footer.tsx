import { Github, BookOpen } from "lucide-react";

const Footer = () => {
  return (
    <footer className="border-t border-border py-8 mt-20">
      <div className="container max-w-6xl mx-auto px-4">
        <div className="flex flex-col md:flex-row items-center justify-between gap-4">
          <div className="text-sm text-muted-foreground">
            © {new Date().getFullYear()} EchoDB
          </div>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
