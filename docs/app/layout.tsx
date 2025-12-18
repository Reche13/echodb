import { RootProvider } from "fumadocs-ui/provider/next";
import "./global.css";
import { Inter, JetBrains_Mono } from "next/font/google";

// const jetbrainsMono = JetBrains_Mono({
//   subsets: ['latin'],
//   variable: '--font-mono',
// });

const inter = Inter({
  subsets: ["latin"],
  weight: ["400", "500", "600", "700"],
});

export default function Layout({ children }: LayoutProps<"/">) {
  return (
    <html lang="en" className={inter.className} suppressHydrationWarning>
      <body className="flex flex-col min-h-screen">
        <RootProvider>{children}</RootProvider>
      </body>
    </html>
  );
}
