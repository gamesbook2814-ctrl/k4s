import { Link } from "react-router-dom";
import { Rocket, PartyPopper, Coffee, Building2 } from "lucide-react";

const icons = {
  Rocket,
  PartyPopper,
  Coffee,
  Building2,
};

const businessNeeds = [
  {
    id: "startups",
    name: "For Startups",
    description: "Essential printing for new businesses",
    icon: "Rocket",
    href: "/shop-by/startups",
  },
  {
    id: "events",
    name: "Events & Promotions",
    description: "Banners, standees, and marketing materials",
    icon: "PartyPopper",
    href: "/shop-by/events",
  },
  {
    id: "cafe",
    name: "Cafe & Restaurant",
    description: "Menus, packaging, and branding essentials",
    icon: "Coffee",
    href: "/shop-by/cafe-restaurant",
  },
  {
    id: "office",
    name: "Office Essentials",
    description: "Stationery and business supplies",
    icon: "Building2",
    href: "/shop-by/office",
  },
];

const BusinessNeeds = () => {
  return (
    <section className="py-12 md:py-16 bg-muted">
      <div className="container">
        <h2 className="text-2xl md:text-3xl font-heading font-bold text-center mb-10">
          Shop By Business Needs
        </h2>
        <div className="grid grid-cols-2 md:grid-cols-4 gap-4 md:gap-6">
          {businessNeeds.map((item) => {
            const Icon = icons[item.icon as keyof typeof icons];
            return (
              <Link
                key={item.id}
                to={item.href}
                className="bg-background rounded-xl p-6 text-center shadow-card hover:shadow-card-hover transition-all hover:-translate-y-1 group"
              >
                <div className="w-14 h-14 mx-auto mb-4 rounded-full bg-primary/10 flex items-center justify-center group-hover:bg-primary/20 transition-colors">
                  <Icon className="w-7 h-7 text-primary" />
                </div>
                <h3 className="font-heading font-semibold text-foreground mb-1">
                  {item.name}
                </h3>
                <p className="text-sm text-muted-foreground">
                  {item.description}
                </p>
              </Link>
            );
          })}
        </div>
      </div>
    </section>
  );
};

export default BusinessNeeds;
