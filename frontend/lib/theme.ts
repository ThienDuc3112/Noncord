// lib/theme.ts
export const theme = {
  colors: {
    // Background gradients / surfaces
    background: {
      // Use with `bg-gradient-to-br`
      primary: "from-[#1e2030] to-[#181926]", // mantle -> crust
      card: "bg-[#24273a]/95", // base
    },

    // Text colors
    text: {
      primary: "text-[#cad3f5]", // text
      secondary: "text-[#a5adcb]", // subtext0
      muted: "text-[#6e738d]", // overlay0
      placeholder: "placeholder:text-[#6e738d]",
    },

    // Interactive elements
    interactive: {
      primary: "bg-[#c6a0f6] hover:bg-[#b7bdf8] text-[#181926]", // mauve -> lavender
      link: "text-[#f5bde6] hover:underline", // rosewater
      focus: "focus-visible:ring-[#c6a0f6]", // mauve
    },

    // Form elements
    form: {
      input: "bg-[#1e2030] border-none", // mantle
      label: "text-[#a5adcb]", // subtext0
      required: "text-[#ed8796]", // red
      separator: "bg-[#363a4f]", // surface0
    },

    // States
    states: {
      error: "text-[#ed8796]", // red
      success: "text-[#a6da95]", // green
      warning: "text-[#eed49f]", // yellow
    },
  },

  // Common class combinations
  classes: {
    input:
      "bg-[#1e2030] border border-transparent text-[#cad3f5] " +
      "placeholder:text-[#6e738d] focus:ring-0 focus:ring-offset-0 " +
      "focus-visible:ring-1 focus-visible:ring-[#c6a0f6]",

    button: {
      primary:
        "bg-[#c6a0f6] hover:bg-[#b7bdf8] text-[#181926] font-medium py-3 rounded-sm",
      secondary:
        "bg-[#cad3f5] hover:bg-[#b8c0e0] text-[#181926] border-none font-medium py-3",
    },

    label: "text-[#a5adcb] text-xs font-bold uppercase tracking-wide",

    card: "bg-[#24273a]/95 border border-[#363a4f] shadow-2xl rounded-md",

    background: "min-h-screen bg-gradient-to-br from-[#1e2030] to-[#181926]",

    formMessage: "text-[#ed8796] text-xs mt-1 font-medium",

    inputError:
      "bg-[#1e2030] border border-[#ed8796] text-[#cad3f5] " +
      "placeholder:text-[#6e738d] focus:ring-0 focus:ring-offset-0 " +
      "focus-visible:ring-1 focus-visible:ring-[#ed8796]",
  },
} as const;

// Background pattern SVG (Macchiato-ish dot pattern)
export const backgroundPattern =
  "url(\"data:image/svg+xml,%3Csvg width='60' height='60' viewBox='0 0 60 60' xmlns='http://www.w3.org/2000/svg'%3E%3Cg fill='none' fillRule='evenodd'%3E%3Cg fill='%23b7bdf8' fillOpacity='0.08'%3E%3Ccircle cx='30' cy='30' r='2'/%3E%3C/g%3E%3C/g%3E%3C/svg%3E\")";
