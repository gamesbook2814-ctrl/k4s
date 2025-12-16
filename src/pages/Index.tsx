import Layout from "@/components/layout/Layout";
import HeroCarousel from "@/components/home/HeroCarousel";
import PromoBanner from "@/components/home/PromoBanner";
import BusinessNeeds from "@/components/home/BusinessNeeds";
import FeaturedCategories from "@/components/home/FeaturedCategories";
import BestSellers from "@/components/home/BestSellers";
import Testimonials from "@/components/home/Testimonials";

const Index = () => {
  return (
    <Layout>
      <HeroCarousel />
      <PromoBanner />
      <BusinessNeeds />
      <FeaturedCategories />
      <BestSellers />
      <Testimonials />
    </Layout>
  );
};

export default Index;
