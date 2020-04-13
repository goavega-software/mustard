export const scriptLoader = async (src: string, async: boolean) => {
  return new Promise<void>((resolve: () => void, reject) => {
    if (document.querySelector('script[src="' + src + '"]')) {
      resolve();
      return;
    }

    const el = document.createElement("script");

    el.type = "text/javascript";
    el.async = async;
    el.src = src;

    el.addEventListener("load", () => resolve());
    el.addEventListener("error", reject);
    el.addEventListener("abort", reject);
    document.head.appendChild(el);
  });
};
