### How to run it?
```docker-compose up -d```

### Test
Enter `http://127.0.0.1:8081/5` (https://github.com/lbarulski/esi-showcase/blob/main/main.go#L17), first load should take over 5s to first part of content (shell content), and slowly loading ESI by esi, calling synchronously delayed fragments (https://github.com/lbarulski/esi-showcase/blob/main/main.go#L24). 
After refresh, second load should be fast with shell content (already cached), but still behaving exactly the same with fragments (ESI includes).

Here (https://github.com/lbarulski/esi-showcase/blob/main/templates/index.tmpl) you can see how that looks from backend<->varnish communication perspective, main endpoint is just returning this HTML, and starts streaming right away, once found `esi:include` tag, makes request to endpoints returning other fragments of HTML (https://github.com/lbarulski/esi-showcase/blob/main/templates/delay.tmpl)
