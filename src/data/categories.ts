export interface SubCategory {
  name: string;
  href: string;
  isNew?: boolean;
}

export interface Category {
  name: string;
  href: string;
  subcategories: SubCategory[];
}

export interface MegaMenuSection {
  title: string;
  items: SubCategory[];
}

export const mainCategories = [
  { name: "All Products", href: "/products" },
  { name: "Same Day Delivery", href: "/same-day-delivery" },
  { name: "Calendars & Diaries", href: "/calendars-diaries" },
  { name: "Packaging", href: "/packaging" },
  { name: "Apparel", href: "/apparel" },
  { name: "Stationery", href: "/stationery" },
  { name: "Corporate Gifts", href: "/corporate-gifts" },
  { name: "Photo Gifts", href: "/photo-gifts" },
  { name: "Drinkware", href: "/drinkware" },
];

export const megaMenuData: Record<string, MegaMenuSection[]> = {
  "All Products": [
    {
      title: "Browse All Categories",
      items: [
        { name: "Calendars & Diaries", href: "/calendars-diaries" },
        { name: "Apparels", href: "/apparel" },
        { name: "Photo Gifts", href: "/photo-gifts" },
        { name: "Rewards and Recognition", href: "/rewards-recognition" },
        { name: "Business Stationery", href: "/stationery" },
        { name: "Corporate Gifts", href: "/corporate-gifts" },
        { name: "Packaging", href: "/packaging" },
        { name: "Marketing & Promo", href: "/marketing-promo" },
        { name: "Premium Products", href: "/premium-products" },
        { name: "Drinkware", href: "/drinkware", isNew: true },
        { name: "Gift Hampers", href: "/gift-hampers" },
        { name: "Sample Kit", href: "/sample-kit" },
      ],
    },
    {
      title: "Best Sellers",
      items: [
        { name: "Visiting Cards", href: "/products/visiting-cards" },
        { name: "Round Neck T-shirts", href: "/products/round-neck-tshirts" },
        { name: "Labels", href: "/products/labels" },
        { name: "Booklets", href: "/products/booklets" },
        { name: "Stickers", href: "/products/stickers" },
        { name: "Brochures", href: "/products/brochures" },
        { name: "Notebooks", href: "/products/notebooks" },
        { name: "Mugs", href: "/products/mugs" },
        { name: "Button Badge", href: "/products/button-badge" },
        { name: "Photo Books", href: "/products/photo-books" },
        { name: "Tote Bags", href: "/products/tote-bags" },
        { name: "Photo Prints", href: "/products/photo-prints" },
      ],
    },
    {
      title: "Business Stationery",
      items: [
        { name: "View All →", href: "/stationery" },
        { name: "Business Cards", href: "/products/business-cards" },
        { name: "Letterheads", href: "/products/letterheads" },
        { name: "Stamps", href: "/products/stamps" },
        { name: "Envelopes", href: "/products/envelopes" },
      ],
    },
    {
      title: "Gift Hampers",
      items: [
        { name: "View All →", href: "/gift-hampers" },
        { name: "Employee Engagement Kits", href: "/products/employee-kits" },
        { name: "Welcome Kits", href: "/products/welcome-kits" },
        { name: "Awards & Trophies", href: "/products/awards-trophies" },
        { name: "Sustainable Kits", href: "/products/sustainable-kits" },
      ],
    },
  ],
  "Calendars & Diaries": [
    {
      title: "Calendars & Diaries",
      items: [
        { name: "View All →", href: "/calendars-diaries" },
        { name: "Everyday Collection", href: "/products/everyday-calendars" },
        { name: "Classic Collection", href: "/products/classic-calendars" },
        { name: "Prestige Collection", href: "/products/prestige-calendars" },
        { name: "Calendars", href: "/products/calendars" },
      ],
    },
    {
      title: "Apparels",
      items: [
        { name: "View All →", href: "/apparel" },
        { name: "Custom Round Neck T-shirts", href: "/products/round-neck-tshirts" },
        { name: "Custom Polo T-shirts", href: "/products/polo-tshirts" },
        { name: "Custom Hoodies", href: "/products/hoodies" },
        { name: "Custom Caps", href: "/products/caps" },
      ],
    },
    {
      title: "Photo Gifts",
      items: [
        { name: "View All →", href: "/photo-gifts" },
        { name: "Photo Frames", href: "/products/photo-frames" },
        { name: "Photo Prints", href: "/products/photo-prints" },
        { name: "Acrylic Photo Prints", href: "/products/acrylic-prints" },
        { name: "PhotoBooks", href: "/products/photo-books", isNew: true },
      ],
    },
    {
      title: "Rewards and Recognition",
      items: [
        { name: "View All →", href: "/rewards-recognition" },
        { name: "Crystal & Acrylic Awards", href: "/products/crystal-awards" },
        { name: "Wooden Trophies & Mementos", href: "/products/wooden-trophies" },
        { name: "Certificates", href: "/products/certificates" },
        { name: "Medals", href: "/products/medals" },
      ],
    },
  ],
  "Packaging": [
    {
      title: "Packaging",
      items: [
        { name: "View All →", href: "/packaging" },
        { name: "Flat Mailer Boxes", href: "/products/mailer-boxes" },
        { name: "Stickers", href: "/products/stickers" },
        { name: "Labels", href: "/products/labels" },
        { name: "Tote Bags", href: "/products/tote-bags", isNew: true },
      ],
    },
    {
      title: "Marketing & Promo",
      items: [
        { name: "View All →", href: "/marketing-promo" },
        { name: "Banners", href: "/products/banners" },
        { name: "Booklets", href: "/products/booklets" },
        { name: "Brochures", href: "/products/brochures" },
        { name: "Standees", href: "/products/standees" },
      ],
    },
    {
      title: "Premium Products",
      items: [
        { name: "View All →", href: "/premium-products" },
        { name: "Business Cards", href: "/products/premium-business-cards" },
        { name: "Cards", href: "/products/premium-cards" },
        { name: "Stickers", href: "/products/premium-stickers" },
        { name: "Notebooks", href: "/products/premium-notebooks" },
      ],
    },
    {
      title: "Drinkware",
      items: [
        { name: "View All →", href: "/drinkware", isNew: true },
        { name: "Best Sellers", href: "/products/drinkware-bestsellers" },
        { name: "Premium Mugs", href: "/products/premium-mugs" },
        { name: "Premium Flasks", href: "/products/premium-flasks" },
        { name: "Premium Sipper Bottles", href: "/products/sipper-bottles" },
      ],
    },
  ],
  "Corporate Gifts": [
    {
      title: "Corporate Gifts",
      items: [
        { name: "View All →", href: "/corporate-gifts" },
        { name: "Drinkware", href: "/drinkware" },
        { name: "Backpack", href: "/products/backpack", isNew: true },
        { name: "Awards & Trophies", href: "/products/awards-trophies" },
        { name: "Employee Engagement Kits", href: "/products/employee-kits" },
      ],
    },
    {
      title: "Sample Kit",
      items: [
        { name: "View All →", href: "/sample-kit" },
        { name: "Flat Mailer Boxes", href: "/products/mailer-boxes" },
        { name: "Stickers", href: "/products/stickers" },
        { name: "Flexible Pouches", href: "/products/flexible-pouches" },
        { name: "Tote Bags", href: "/products/tote-bags" },
      ],
    },
    {
      title: "Shop By Needs",
      items: [
        { name: "For Startups", href: "/shop-by/startups" },
        { name: "Events & Promotions", href: "/shop-by/events" },
        { name: "Cafe & Restaurant Essentials", href: "/shop-by/cafe-restaurant" },
        { name: "Employee Engagement Kits", href: "/products/employee-kits" },
      ],
    },
    {
      title: "Stores & Services",
      items: [
        { name: "Same Day Delivery", href: "/same-day-delivery" },
        { name: "Bulk Buying", href: "/bulk-buying" },
        { name: "Store Locator", href: "/store-locator" },
        { name: "New Launches", href: "/new-launches" },
        { name: "Blog", href: "/blog" },
      ],
    },
  ],
};

export const featuredCategories = [
  {
    id: "calendars",
    name: "Calendars & Diaries",
    description: "Premium quality calendars and diaries for 2026",
    image: "/placeholder.svg",
    href: "/calendars-diaries",
  },
  {
    id: "business-cards",
    name: "Business Cards",
    description: "Professional visiting cards for your business",
    image: "/placeholder.svg",
    href: "/products/business-cards",
  },
  {
    id: "tshirts",
    name: "Custom T-Shirts",
    description: "High quality custom printed apparel",
    image: "/placeholder.svg",
    href: "/apparel",
  },
  {
    id: "packaging",
    name: "Packaging",
    description: "Custom boxes, labels, and packaging solutions",
    image: "/placeholder.svg",
    href: "/packaging",
  },
  {
    id: "photo-gifts",
    name: "Photo Gifts",
    description: "Personalized photo frames, mugs, and more",
    image: "/placeholder.svg",
    href: "/photo-gifts",
  },
  {
    id: "corporate-gifts",
    name: "Corporate Gifts",
    description: "Premium gifts for your business partners",
    image: "/placeholder.svg",
    href: "/corporate-gifts",
  },
];

export const businessNeeds = [
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
