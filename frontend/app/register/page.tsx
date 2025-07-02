"use client";

import type React from "react";

import { useState } from "react";
import Link from "next/link";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Eye, EyeOff } from "lucide-react";
import { AuthLayout } from "@/components/auth-layout";
import { theme } from "@/lib/theme";

export default function RegisterPage() {
  const [showPassword, setShowPassword] = useState(false);
  const [formData, setFormData] = useState({
    email: "",
    username: "",
    password: "",
  });

  const handleInputChange = (field: string, value: string) => {
    setFormData((prev) => ({ ...prev, [field]: value }));
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    console.log("Register submitted:", formData);
  };

  return (
    <AuthLayout>
      <div className="text-center mb-6">
        <h1 className={`text-2xl font-bold ${theme.colors.text.primary} mb-2`}>
          Create an account
        </h1>
        <p className={`${theme.colors.text.secondary} text-sm`}>
          Join our amazing community!
        </p>
      </div>

      <form onSubmit={handleSubmit} className="space-y-4">
        <div className="space-y-2">
          <Label htmlFor="email" className={theme.classes.label}>
            Email <span className={theme.colors.states.error}>*</span>
          </Label>
          <Input
            id="email"
            type="email"
            value={formData.email}
            onChange={(e) => handleInputChange("email", e.target.value)}
            className={theme.classes.input}
            required
          />
        </div>

        <div className="space-y-2">
          <Label htmlFor="username" className={theme.classes.label}>
            Username <span className={theme.colors.states.error}>*</span>
          </Label>
          <Input
            id="username"
            type="text"
            value={formData.username}
            onChange={(e) => handleInputChange("username", e.target.value)}
            className={theme.classes.input}
            required
          />
        </div>

        <div className="space-y-2">
          <Label htmlFor="password" className={theme.classes.label}>
            Password <span className={theme.colors.states.error}>*</span>
          </Label>
          <div className="relative">
            <Input
              id="password"
              type={showPassword ? "text" : "password"}
              value={formData.password}
              onChange={(e) => handleInputChange("password", e.target.value)}
              className={theme.classes.input + " pr-10"}
              required
            />
            <button
              type="button"
              onClick={() => setShowPassword(!showPassword)}
              className={`absolute right-3 top-1/2 -translate-y-1/2 ${theme.colors.text.secondary} hover:text-white`}
            >
              {showPassword ? (
                <EyeOff className="h-4 w-4" />
              ) : (
                <Eye className="h-4 w-4" />
              )}
            </button>
          </div>
        </div>

        <Button
          type="submit"
          className={`w-full ${theme.classes.button.primary}`}
        >
          Continue
        </Button>

        <p className={`${theme.colors.text.muted} text-xs leading-relaxed`}>
          By registering, you agree to Discord's{" "}
          <button className={theme.colors.interactive.link}>
            Terms of Service
          </button>{" "}
          and{" "}
          <button className={theme.colors.interactive.link}>
            Privacy Policy
          </button>
          .
        </p>

        <Link
          href="/login"
          className={`${theme.colors.interactive.link} text-sm w-full text-center block`}
        >
          Already have an account?
        </Link>
      </form>
    </AuthLayout>
  );
}
