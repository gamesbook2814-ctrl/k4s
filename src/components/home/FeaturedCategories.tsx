import { Link } from "react-router-dom";
import { Calendar, CreditCard, Shirt, Package, Image, Gift } from "lucide-react";

const categories = [
  {
    id: "calendars",
    name: "Calendars & Diaries",
    description: "Premium quality calendars and diaries for 2026",
    icon: Calendar,
    href: "/calendars-diaries",
    color: "bg-violet-500",
  },
  {
    id: "business-cards",
    name: "Business Cards",
    description: "Professional visiting cards for your business",
    icon: CreditCard,
    href: "/products/business-cards",
    color: "bg-blue-500",
  },
  {
    id: "tshirts",
    name: "Custom T-Shirts",
    description: "High quality custom printed apparel",
    icon: Shirt,
    href: "/apparel",
    color: "bg-emerald-500",
  },
  {
    id: "packaging",
    name: "Packaging",
    description: "Custom boxes, labels, and packaging solutions",
    icon: Package,
    href: "/packaging",
    color: "bg-amber-500",
  },
  {
    id: "photo-gifts",
    name: "Photo Gifts",
    description: "Personalized photo frames, mugs, and more",
    icon: Image,
    href: "/photo-gifts",
    color: "bg-pink-500",
  },
  {
    id: "corporate-gifts",
    name: "Corporate Gifts",
    description: "Premium gifts for your business partners",
    icon: Gift,
    href: "/corporate-gifts",
    color: "bg-indigo-500",
  },
];

const FeaturedCategories = () => {
  return (
    <section className="py-12 md:py-16">
      <div className="container">
        <h2 className="text-2xl md:text-3xl font-heading font-bold text-center mb-3">
          Shop By Category
        </h2>
        <p className="text-muted-foreground text-center mb-10 max-w-2xl mx-auto">
          Explore our wide range of printing products and services
        </p>
        <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gap-4 md:gap-6">
          {categories.map((category) => {
            const Icon = category.icon;
            return (
              <Link
                key={category.id}
                to={category.href}
                className="bg-card rounded-xl p-5 text-center border border-border hover:border-primary/50 hover:shadow-lg transition-all group"
              >
                <div className={`w-16 h-16 mx-auto mb-4 rounded-2xl ${category.color} flex items-center justify-center group-hover:scale-110 transition-transform`}>
                  <Icon className="w-8 h-8 text-white" />
                </div>
                <h3 className="font-heading font-semibold text-foreground text-sm">
                  {category.name}
                </h3>
              </Link>
            );
          })}
        </div>
      </div>
    </section>
  );
};

export default FeaturedCategories;
