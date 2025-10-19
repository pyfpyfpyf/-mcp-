package main

import (
	"context"
	"fmt"
	"html"
	"strings"
)

// GenerateNanoBananaCoverPrompt 返回“Nano Banana 一键稳定生成小红书封面图”的完整提示词模板。
// 如果提供了 topic，将自动填充到“用户输入区”的话题占位中。
func (s *XiaohongshuService) GenerateNanoBananaCoverPrompt(topic string) string {
	var b strings.Builder
	b.WriteString("Nano Banana 一键稳定生成小红书封面图提示词分享\n\n")
	b.WriteString("这套方案的设计亮点：\n\n")
	b.WriteString("图文分离，各司其职：因为 Nano Banana 的中文不行，就让 Nano Banana 专心生成高质量插画图，再用 HTML 代码把文字精准地“贴”上去。AI的创意和代码的稳定，两不耽误。而且这套方案不仅支持 Nano Banana，还支持所有的绘图模型，上面6宫格里，下面三个就是 Flux kontext 的图。\n\n")
	b.WriteString("告别文字抽卡玄学：彻底解决AI写字位置随机、样式丑的痛点。标题位置、字体、大小全用代码写死，保证每张图都专业统一。\n\n")
	b.WriteString("模板化，一键复用：版式一次搞定，就成了万能模板。以后换个话题，只需改下要求，风格全部统一，非常适合打造个人品牌调性。\n\n")
	b.WriteString("方案使用：\n\n")
	b.WriteString("这套方案最方便的是在 Lovart 这种设计 Agent 里使用，只需要把提示词复制粘贴进去，改个话题，一键生成。\n\n")
	b.WriteString("如果不用 Lovart ，也可以通过其他 Agent 或 Chatwise、Claude Code 等产品分别绘制图像和HTML。\n\n")
	b.WriteString("提示词：注意，请自行修改风格、话题。\n")
	b.WriteString("### **Prompt: 热点事件话题卡片生成器 (v1.0)**\n\n")
	b.WriteString("#### **🎯 核心任务**\n\n")
	b.WriteString("针对一个给定的**热点事件或话题**，创建一个 **宽600x高800 像素**的视觉卡片。最终产出物是一个**可直接在浏览器中打开的、包含内联CSS的完整HTML文件**。\n\n")
	b.WriteString("#### **🎨 视觉风格总览**\n\n")
	b.WriteString("*   **主题风格**: 现代风格化插画 (Modern, Stylized, Illustrative)。风格需与事件调性匹配，可以是概念性的、象征性的或场景再现的，包含一些英文的关键词。\n")
	b.WriteString("*   **色彩**: **与事件主题和情感基调相匹配**。色彩需干净、有冲击力但不过度饱和，构图清晰，主次分明。\n")
	b.WriteString("*   **构图**: 简洁、不拥挤，通过视觉语言精准传达事件核心。\n\n")
	b.WriteString("---\n\n")
	b.WriteString("### **📐 页面布局与规范**\n\n")
	b.WriteString("#### **第一部分：插图区 (上半部分, 约占70%高度)**\n\n")
	b.WriteString("此区域用于展示高度概念化、与事件相关的风格化插图。\n\n")
	b.WriteString("*   **尺寸与布局**:\n")
	b.WriteString("    * **图片比例：1:1**  \n")
	b.WriteString("    * **图片模型（严格执行）：nano banana**\n")
	b.WriteString("    * **满幅贴边 (Full-bleed)**：插图必须完全覆盖此区域，无任何内边距或留白。\n")
	b.WriteString("    *   **绝对直角**: 四个角必须是 **90° 直角**。\n")
	b.WriteString("*   **【极其重要】禁止的视觉元素**:\n")
	b.WriteString("    *   **严禁任何形式的边框**: 包括但不限于外框、相框、白边、贴纸边、描边围框、UI面板、容器、发光边、投影等。\n")
	b.WriteString("    *   **严禁圆角**: 整个插图区域必须是完美的矩形。\n")
	b.WriteString("*   **插图生成要求**:\n")
	b.WriteString("    *   **内容**: 插图需要**象征性或概念性地表现事件的核心**，而非简单复刻新闻照片。应引发思考。\n")
	b.WriteString("    *   **风格**: 现代、简洁、风格化的插画，可以带有轻微的肌理感或扁平化风格。\n")
	b.WriteString("    *   **光影与构图**: 构图具有设计感 (`clean composition`)，视觉焦点明确，避免元素堆砌导致画面混乱 (`avoid overcrowding`)。\n\n")
	b.WriteString("---\n\n")
	b.WriteString("#### **第二部分：信息区 (下半部分, 约占30%高度)**\n\n")
	b.WriteString("此区域用于展示对事件的提炼和点评，背景为纯色。\n\n")
	b.WriteString("*   **布局**: 所有文本内容**水平居中对齐**，垂直方向上舒适地分布。\n")
	b.WriteString("*   **背景**: 纯净的单色背景通常为黑色、白色、浅灰，需要与插图的主色协调，**禁止使用卡片、阴影、边框**等任何装饰。\n")
	b.WriteString("*   **内容与样式**:\n")
	b.WriteString("    1.  **核心标题 (Headline)**:\n")
	b.WriteString("        *   内容: 一句话精准概括事件的核心或观点。\n")
	b.WriteString("        *   样式: **加粗**, 字号 `34px`, 颜色 `#111`\n")
	b.WriteString("    2.  **引言/评论 (Quote)**:\n")
	b.WriteString("        *   内容: 一段相关的引言、数据或一句发人深省的评论。\n")
	b.WriteString("        *   样式: 字号 `21px`, 颜色 `#555`, 行高 `1.6`\n\n")
	b.WriteString("---\n\n")
	b.WriteString("### **💡 系统执行流程**\n\n")
	b.WriteString("1.  接收用户输入的**热点事件关键词**。\n")
	b.WriteString("2.  **分析与提炼**:\n")
	b.WriteString("    *   理解事件的核心内容和情感基调。\n")
	b.WriteString("    *   构思一句强有力的**核心标题**。\n")
	b.WriteString("    *   寻找或创作一句精辟的**引言/评论**。\n")
	b.WriteString("3.  **生成插图**:\n")
	b.WriteString("    *   根据对事件的理解，调用图像生成模型，创作一幅符合要求的概念插图。\n")
	b.WriteString("    *   严格遵守【插图区】的所有视觉规范，特别是**满幅、直角、无边框**的要求。\n")
	b.WriteString("4.  **构建HTML**:\n")
	b.WriteString("    *   生成一个完整的HTML文件结构 (`<html>`, `<head>`, `<body>`)。\n")
	b.WriteString("    *   在 `<head>` 中嵌入 `<style>` 标签，包含所有内联CSS。\n")
	b.WriteString("    *   **字体**: 优先使用现代无衬线字体，如 `'Inter'`, `'Noto Sans SC'`, 或系统默认字体。\n")
	b.WriteString("    *   **布局CSS**:\n")
	b.WriteString("        *   `body` 设置 `margin: 0;`\n")
	b.WriteString("        *   上半部分容器 `height: 70vh;`，背景图 `background-size: cover; background-position: center;`\n")
	b.WriteString("        *   下半部分容器 `height: 30vh;`，使用 Flexbox 实现内容的垂直和水平居中，并设置合适的 `padding`。\n")
	b.WriteString("    *   将生成的插图作为上半部分的背景图像。\n")
	b.WriteString("    *   将核心标题和引言填入下半部分，并应用指定的CSS样式。\n")
	b.WriteString("5.  **输出**: 返回最终的、单一的HTML文件。\n\n")
	b.WriteString("---\n\n")
	b.WriteString("### **✍️ 用户输入区**\n\n")
	b.WriteString("请在下方简要描述一个热点事件或话题：\n\n")
	b.WriteString("``" + "`\n")
	b.WriteString("话题：【在此填写事件或话题，可以直接粘贴全文进来】\n")
	b.WriteString("``" + "`\n")
	b.WriteString("**示例输入**: `全球芯片供应链紧张`, `人工智能对创意产业的影响`, `火星探测的新发现`\n\n")
	b.WriteString("好的，今天就是有机大橘子为大家带来的 Nano Banana 教程。\n")
	b.WriteString("希望大家玩得开心，我们下次再见。")

	tpl := b.String()
	if strings.TrimSpace(topic) != "" {
		return strings.Replace(tpl, "话题：【在此填写事件或话题，可以直接粘贴全文进来】", fmt.Sprintf("话题：【%s】", topic), 1)
	}
	return tpl
}

// GenerateCoverHTML 生成符合规范的小红书封面 HTML（600x800，插图上70%，信息区下30%）。
func (s *XiaohongshuService) GenerateCoverHTML(_ context.Context, imageURL, headline, quote, background string) (string, error) {
	img := strings.TrimSpace(imageURL)
	h := html.EscapeString(strings.TrimSpace(headline))
	q := html.EscapeString(strings.TrimSpace(quote))
	bg := strings.TrimSpace(background)
	if bg == "" {
		bg = "#fff"
	}
	// 简单校验
	if img == "" || h == "" || q == "" {
		return "", fmt.Errorf("image_url、headline、quote 不能为空")
	}

	htmlDoc := fmt.Sprintf(`<!DOCTYPE html>
<html lang="zh-CN">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>小红书封面</title>
  <style>
    :root { --bg: %s; }
    * { box-sizing: border-box; }
    html, body { margin: 0; padding: 0; }
    body {
      margin: 0;
      background: #f5f5f5;
      font-family: "Inter", "Noto Sans SC", system-ui, -apple-system, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol";
      display: flex;
      align-items: center;
      justify-content: center;
      min-height: 100vh;
    }
    .card {
      width: 600px;
      height: 800px;
      background: #fff;
      display: flex;
      flex-direction: column;
      border: none;
      border-radius: 0;
      overflow: hidden;
    }
    .illustration {
      flex: 7 0 0; /* 70%% */
      background-image: url('%s');
      background-size: cover;
      background-position: center center;
      background-repeat: no-repeat;
      margin: 0;
      border: none;
      border-radius: 0; /* 绝对直角 */
    }
    .info {
      flex: 3 0 0; /* 30%% */
      background: var(--bg);
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      padding: 18px 24px;
      text-align: center;
      gap: 10px;
    }
    .headline {
      margin: 0;
      font-size: 34px;
      font-weight: 800;
      color: #111;
      letter-spacing: 0.3px;
    }
    .quote {
      margin: 0;
      font-size: 21px;
      color: #555;
      line-height: 1.6;
      font-weight: 500;
    }
  </style>
</head>
<body>
  <div class="card">
    <div class="illustration"></div>
    <div class="info">
      <h1 class="headline">%s</h1>
      <p class="quote">%s</p>
    </div>
  </div>
</body>
</html>`, bg, img, h, q)

	return htmlDoc, nil
}
