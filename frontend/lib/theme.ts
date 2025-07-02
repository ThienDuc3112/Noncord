export const theme = {
  colors: {
    // Background gradients
    background: {
      primary: "from-pink-900 to-rose-900",
      card: "bg-gray-900/95",
    },

    // Text colors
    text: {
      primary: "text-white",
      secondary: "text-[#b9bbbe]",
      muted: "text-[#72767d]",
      placeholder: "placeholder:text-[#72767d]",
    },

    // Interactive elements
    interactive: {
      primary: "bg-pink-700 hover:bg-pink-800",
      link: "text-pink-300 hover:underline",
      focus: "focus-visible:ring-pink-600",
    },

    // Form elements
    form: {
      input: "bg-[#202225] border-none",
      label: "text-[#b9bbbe]",
      required: "text-red-400",
      separator: "bg-[#4f545c]",
    },

    // States
    states: {
      error: "text-red-400",
      success: "text-green-400",
      warning: "text-yellow-400",
    },
  },

  // Common class combinations
  classes: {
    input:
      "bg-[#202225] border-none text-white placeholder:text-[#72767d] focus:ring-0 focus:ring-offset-0 focus-visible:ring-1 focus-visible:ring-pink-600",
    button: {
      primary:
        "bg-pink-700 hover:bg-pink-800 text-white font-medium py-3 rounded-sm",
      secondary:
        "bg-white hover:bg-gray-100 text-black border-none font-medium py-3",
    },
    label: "text-[#b9bbbe] text-xs font-bold uppercase tracking-wide",
    card: "bg-gray-900/95 border-none shadow-2xl",
    background: "min-h-screen bg-gradient-to-br from-pink-900 to-rose-900",
  },
} as const;

// Background pattern SVG
export const backgroundPattern =
  "url(\"data:image/svg+xml,%3Csvg width='60' height='60' viewBox='0 0 60 60' xmlns='http://www.w3.org/2000/svg'%3E%3Cg fill='none' fillRule='evenodd'%3E%3Cg fill='%23fce7f3' fillOpacity='0.1'%3E%3Ccircle cx='30' cy='30' r='2'/%3E%3C/g%3E%3C/g%3E%3C/svg%3E\")";
