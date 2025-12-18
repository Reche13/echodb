import type { BaseLayoutProps } from "fumadocs-ui/layouts/shared";

export function baseOptions(): BaseLayoutProps {
  return {
    nav: {
      title: "Echo DB",
      transparentMode: "always",
    },
    githubUrl: "https://github.com/reche13/echodb",
  };
}
