{
  "$schema": "https://mintlify.com/schema.json",
  "name": "Starter Kit",
  "logo": {
    "dark": "/logo/dark.svg",
    "light": "/logo/light.svg"
  },
  "favicon": "/favicon.png",
  "colors": {
    "primary": "#0D9373",
    "light": "#07C983",
    "dark": "#0D9373",
    "anchors": {
      "from": "#0D9373",
      "to": "#07C983"
    }
  },
  "feedback": {
    "thumbsRating": true,
    "suggestEdit": true,
    "raiseIssue": true
  },
  "modeToggle": {
    "default": "dark"
  },
  "analytics": {
    "ga4": {
      "measurementId": "G-S98KMZZZY1"
    }
  },
  "topbarLinks": [
    {
      "name": "Learn More",
      "url": "https://www.benzinga.com/apis"
    }
  ],
  "topbarCtaButton": {
    "name": "Get Started",
    "url": "https://www.benzinga.com/apis/licensing/register"
  },
  "api": {
    "auth": {
      "method": "key"
    },
    "playground": {
      "mode": "show"
    }
  },
  "anchors": [
    {
      "name": "API Suite",
      "icon": "list",
      "url": "https://www.benzinga.com/apis/data"
    },
    {
      "name": "Blog",
      "icon": "newspaper",
      "url": "https://www.benzinga.com/apis/blog/"
    },
    {
      "name": "Contact Us",
      "icon": "envelope",
      "url": "mailto:licensing@benzinga.com"
    }
  ],
  "navigation": [
    {
      "group": "Introduction",
      "pages": [
        "home",
        "introduction"
      ]
    },
    {
      "group": "Benzinga APIs",
      "pages": [
        "benzinga-apis/introduction",
        {{REPLACE_HERE}}
      ]
    }
  ],
  "footerSocials": {
    "twitter": "https://twitter.com/Benzinga",
    "facebook": "https://www.facebook.com/Benzinga/",
    "website": "https://www.benzinga.com/apis/",
    "github": "https://github.com/benzinga",
    "linkedin": "https://www.linkedin.com/company/benzinga/"
  }
}