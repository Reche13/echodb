import HeroSection from "@/components/sections/HeroSection";
import FeaturesSection from "@/components/sections/FeaturesSection";
import QuickStartSection from "@/components/sections/QuickStartSection";
import InstallSection from "@/components/sections/InstallSection";
import CTASection from "@/components/sections/CTASection";
import Footer from "@/components/Footer";

export default function HomePage() {
  return (
    <main>
      <HeroSection />
      <FeaturesSection />
      <QuickStartSection />
      <InstallSection />
      <CTASection />
      <Footer />
    </main>
  );
}
