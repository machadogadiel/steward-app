<script setup lang="ts">
import { cva, type VariantProps } from "class-variance-authority";

const button = cva("button", {
  variants: {
    type: {
      primary: "primary",
      secondary: "secondary",
    },
    size: {
      small: "small",
      medium: "medium",
    },
  },
  compoundVariants: [
    { type: "primary", size: "medium", class: "primaryMedium" },
  ],
});

type ButtonProps = VariantProps<typeof button>;

withDefaults(
  defineProps<{ type?: ButtonProps["type"]; size?: ButtonProps["size"] }>(),
  {
    type: "primary",
    size: "medium",
  }
);
</script>

<template>
  <button :class="button({ type, size })">
    <slot />
  </button>
</template>

<style scoped>
.button {
  display: inline-flex;
  border-width: 1px;
  border-radius: 6px;
  border-style: solid;
}

.primary {
  color: rgb(255 255 255);
  background-color: rgb(33, 33, 34);
  border: transparent;
}

.primary:hover {
  background-color: rgb(93, 93, 93);
  
}

.secondary {
  background-color: rgb(255 255 255);
  color: rgb(31 41 55);
  border-color: rgb(156 163 175);
}

.secondary:hover {
  background-color: rgb(243 244 246);
}

.small {
  font-size: 0.875rem /* 14px */;
  line-height: 1.25rem /* 20px */;
  padding: 0.25rem 0.5rem;
}

.medium {
  font-size: 1rem /* 16px */;
  line-height: 1.5rem /* 24px */;
  padding: 0.5rem 1rem;
}


</style>
