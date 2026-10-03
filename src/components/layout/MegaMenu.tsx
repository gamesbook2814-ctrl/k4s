import { useState } from "react";
import { Link } from "react-router-dom";
import { ChevronDown } from "lucide-react";
import { mainCategories, allProductsMegaMenu, categoryMegaMenus } from "@/data/categories";
import { cn } from "@/lib/utils";

const MegaMenu = () => {
  const [activeCategory, setActiveCategory] = useState<string | null>(null);

  return (
    <nav className="bg-background border-b border-border relative">
      <div className="container">
        <ul className="flex items-center gap-1 overflow-x-auto scrollbar-hide py-1">
          {mainCategories.map((category) => {
            const hasMenu = category.name === "All Products" || categoryMegaMenus[category.name];
            return (
              <li
                key={category.name}
                className="relative"
                onMouseEnter={() => setActiveCategory(category.name)}
                onMouseLeave={() => setActiveCategory(null)}
              >
                <Link
                  to={category.href}
                  className={cn(
                    "flex items-center gap-1 px-3 py-3 text-sm font-medium whitespace-nowrap transition-colors",
                    activeCategory === category.name
                      ? "text-primary"
                      : "text-foreground hover:text-primary"
                  )}
                >
                  {category.name}
                  {hasMenu && (
                    <ChevronDown className={cn(
                      "w-4 h-4 transition-transform",
                      activeCategory === category.name && "rotate-180"
                    )} />
                  )}
                </Link>

                {/* All Products Mega Menu - Full Width */}
                {category.name === "All Products" && activeCategory === "All Products" && (
                  <div 
                    className="fixed left-0 right-0 top-auto bg-background border-t border-b border-border shadow-lg animate-slide-down z-50"
                    style={{ marginTop: "1px" }}
                  >
                    <div className="container py-8">
                      <div className="grid grid-cols-6 gap-8">
                        {allProductsMegaMenu.slice(0, 6).map((section, idx) => (
                          <div key={idx}>
                            <h3 className="font-semibold text-primary mb-3 text-sm">
                              {section.title}
                            </h3>
                            <ul className="space-y-2">
                              {section.items.map((item, itemIdx) => (
                                <li key={itemIdx}>
                                  <Link
                                    to={item.href}
                                    className="text-sm text-muted-foreground hover:text-primary transition-colors flex items-center gap-2"
                                  >
                                    {item.name}
                                    {item.isNew && (
                                      <span className="bg-success text-success-foreground text-xs px-2 py-0.5 rounded font-medium">
                                        New
                                      </span>
                                    )}
                                  </Link>
                                </li>
                              ))}
                            </ul>
                          </div>
                        ))}
                      </div>
                      <div className="grid grid-cols-6 gap-8 mt-8 pt-6 border-t border-border">
                        {allProductsMegaMenu.slice(6, 12).map((section, idx) => (
                          <div key={idx}>
                            <h3 className="font-semibold text-primary mb-3 text-sm">
                              {section.title}
                            </h3>
                            <ul className="space-y-2">
                              {section.items.map((item, itemIdx) => (
                                <li key={itemIdx}>
                                  <Link
                                    to={item.href}
                                    className="text-sm text-muted-foreground hover:text-primary transition-colors flex items-center gap-2"
                                  >
                                    {item.name}
                                    {item.isNew && (
                                      <span className="bg-success text-success-foreground text-xs px-2 py-0.5 rounded font-medium">
                                        New
                                      </span>
                                    )}
                                  </Link>
                                </li>
                              ))}
                            </ul>
                          </div>
                        ))}
                      </div>
                      <div className="grid grid-cols-6 gap-8 mt-8 pt-6 border-t border-border">
                        {allProductsMegaMenu.slice(12, 16).map((section, idx) => (
                          <div key={idx}>
                            <h3 className="font-semibold text-primary mb-3 text-sm">
                              {section.title}
                            </h3>
                            <ul className="space-y-2">
                              {section.items.map((item, itemIdx) => (
                                <li key={itemIdx}>
                                  <Link
                                    to={item.href}
                                    className="text-sm text-muted-foreground hover:text-primary transition-colors flex items-center gap-2"
                                  >
                                    {item.name}
                                    {item.isNew && (
                                      <span className="bg-success text-success-foreground text-xs px-2 py-0.5 rounded font-medium">
                                        New
                                      </span>
                                    )}
                                  </Link>
                                </li>
                              ))}
                            </ul>
                          </div>
                        ))}
                      </div>
                    </div>
                  </div>
                )}

                {/* Category Specific Mega Menu */}
                {categoryMegaMenus[category.name] && activeCategory === category.name && (
                  <div 
                    className="absolute left-0 top-full bg-background border border-border shadow-lg rounded-b-lg animate-slide-down z-50"
                    style={{ minWidth: "600px" }}
                  >
                    <div className="grid grid-cols-3 gap-6 p-6">
                      {categoryMegaMenus[category.name].map((section, idx) => (
                        <div key={idx}>
                          <h3 className="font-semibold text-primary mb-3 text-sm">
                            {section.title}
                          </h3>
                          <ul className="space-y-2">
                            {section.items.map((item, itemIdx) => (
                              <li key={itemIdx}>
                                <Link
                                  to={item.href}
                                  className="text-sm text-muted-foreground hover:text-primary transition-colors flex items-center gap-2"
                                >
                                  {item.name}
                                  {item.isNew && (
                                    <span className="bg-success text-success-foreground text-xs px-2 py-0.5 rounded font-medium">
                                      New
                                    </span>
                                  )}
                                </Link>
                              </li>
                            ))}
                          </ul>
                        </div>
                      ))}
                    </div>
                  </div>
                )}
              </li>
            );
          })}
        </ul>
      </div>
    </nav>
  );
};

export default MegaMenu;
