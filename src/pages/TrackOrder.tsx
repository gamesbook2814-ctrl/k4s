import { useState } from "react";
import Layout from "@/components/layout/Layout";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Package, Truck, CheckCircle, Clock } from "lucide-react";

const TrackOrder = () => {
  const [orderId, setOrderId] = useState("");
  const [orderStatus, setOrderStatus] = useState<any>(null);

  const handleTrack = (e: React.FormEvent) => {
    e.preventDefault();
    // Mock order status
    setOrderStatus({
      id: orderId,
      status: "In Transit",
      estimatedDelivery: "Dec 18, 2024",
      steps: [
        { label: "Order Placed", completed: true, date: "Dec 15, 2024" },
        { label: "Processing", completed: true, date: "Dec 15, 2024" },
        { label: "Shipped", completed: true, date: "Dec 16, 2024" },
        { label: "Out for Delivery", completed: false, date: "" },
        { label: "Delivered", completed: false, date: "" },
      ],
    });
  };

  return (
    <Layout>
      <div className="container py-8">
        <h1 className="text-3xl font-heading font-bold mb-8 text-center">Track Your Order</h1>

        <div className="max-w-md mx-auto mb-12">
          <form onSubmit={handleTrack} className="flex gap-3">
            <Input
              placeholder="Enter Order ID (e.g., LK123456)"
              value={orderId}
              onChange={(e) => setOrderId(e.target.value)}
              required
            />
            <Button variant="accent">Track</Button>
          </form>
        </div>

        {orderStatus && (
          <div className="max-w-2xl mx-auto">
            <div className="bg-card rounded-xl border border-border p-6 mb-6">
              <div className="flex justify-between items-start mb-6">
                <div>
                  <p className="text-sm text-muted-foreground">Order ID</p>
                  <p className="font-semibold">{orderStatus.id}</p>
                </div>
                <div className="text-right">
                  <p className="text-sm text-muted-foreground">Estimated Delivery</p>
                  <p className="font-semibold">{orderStatus.estimatedDelivery}</p>
                </div>
              </div>

              <div className="relative">
                {orderStatus.steps.map((step: any, index: number) => (
                  <div key={index} className="flex gap-4 mb-6 last:mb-0">
                    <div className="relative">
                      <div className={`w-8 h-8 rounded-full flex items-center justify-center ${
                        step.completed ? "bg-success text-success-foreground" : "bg-muted text-muted-foreground"
                      }`}>
                        {step.completed ? (
                          <CheckCircle className="w-5 h-5" />
                        ) : (
                          <Clock className="w-5 h-5" />
                        )}
                      </div>
                      {index < orderStatus.steps.length - 1 && (
                        <div className={`absolute left-1/2 top-8 w-0.5 h-8 -translate-x-1/2 ${
                          step.completed ? "bg-success" : "bg-muted"
                        }`} />
                      )}
                    </div>
                    <div className="flex-1 pb-2">
                      <p className={`font-medium ${step.completed ? "text-foreground" : "text-muted-foreground"}`}>
                        {step.label}
                      </p>
                      {step.date && (
                        <p className="text-sm text-muted-foreground">{step.date}</p>
                      )}
                    </div>
                  </div>
                ))}
              </div>
            </div>
          </div>
        )}

        {!orderStatus && (
          <div className="text-center text-muted-foreground py-12">
            <Package className="w-16 h-16 mx-auto mb-4 opacity-50" />
            <p>Enter your order ID to track your order</p>
          </div>
        )}
      </div>
    </Layout>
  );
};

export default TrackOrder;
