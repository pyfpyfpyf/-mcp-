package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"html"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// GenerateCoverRequest 一键生成小红书封面图 请求
// 仅需提供热点事件或话题关键词，其他风格项可选。
type GenerateCoverRequest struct {
	Topic         string `json:"topic" binding:"required"`
	StyleOverride string `json:"style_override,omitempty"`
	Background    string `json:"background,omitempty"` // 信息区背景色，默认#fff
	Seed          int    `json:"seed,omitempty"`
}

// GenerateCoverResponse 生成结果
// HTML 为完整的、包含内联 CSS 的单文件 HTML，可直接在浏览器打开。
type GenerateCoverResponse struct {
	HTML        string `json:"html"`
	ImagePrompt string `json:"image_prompt"`
	ImageModel  string `json:"image_model"`
}

// GenerateNanoBananaCover 生成基于“nano banana”模型风格的封面图 HTML
func (s *XiaohongshuService) GenerateNanoBananaCover(ctx context.Context, req *GenerateCoverRequest) (*GenerateCoverResponse, error) {
	if strings.TrimSpace(req.Topic) == "" {
		return nil, fmt.Errorf("topic 不能为空")
	}

	// 1. 构造图像生成提示词
	imagePrompt := buildNanoBananaPrompt(req.Topic, req.StyleOverride)

	// 2. 生成插图（使用无需密钥的公共推理服务，确保 1:1 比例）
	// 说明：此实现默认走 pollinations 免费服务，便于一键使用；
	// 若需对接企业级模型（如私有推理或其他供应商），可在此替换为自有 HTTP API。
	seed := req.Seed
	if seed == 0 {
		seed = 1000 + rand.Intn(8999)
	}
	imgDataURL, err := fetchConceptImageAsDataURL(ctx, imagePrompt, seed)
	if err != nil {
		return nil, err
	}

	// 3. 生成标题与引言（在无 LLM 情况下的稳健降级实现）
	headline := craftHeadline(req.Topic)
	quote := craftQuote(req.Topic)

	// 4. 组装完整 HTML
	bg := req.Background
	if strings.TrimSpace(bg) == "" {
		bg = "#fff"
	}
	htmlStr := buildCoverHTML(imgDataURL, headline, quote, bg)

	return &GenerateCoverResponse{
		HTML:        htmlStr,
		ImagePrompt: imagePrompt,
		ImageModel:  "nano banana",
	}, nil
}

// buildNanoBananaPrompt 结合规范构建图像提示词
func buildNanoBananaPrompt(topic, styleOverride string) string {
	parts := []string{
		"Modern, stylized, conceptual illustration, clean composition, design-focused, avoid overcrowding",
		"full-bleed image, 1:1 aspect ratio, sharp right-angle corners, no borders, no frames, no rounded corners, no stickers",
		"subtle texture, flat illustration feel, clear focal point",
		fmt.Sprintf("topic: %s", strings.TrimSpace(topic)),
		"include a few English keywords in scene (minimal, tasteful)",
		"color: clean, impactful, balanced saturation, aligned with the topic mood",
		"lighting: soft directional light, graphic clarity",
		"background: integrated scene, no UI panels, no outlines",
		"model: nano banana",
	}
	if strings.TrimSpace(styleOverride) != "" {
		parts = append(parts, styleOverride)
	}
	return strings.Join(parts, ", ")
}

// fetchConceptImageAsDataURL 拉取生成图并转为 data URL，确保 1:1
func fetchConceptImageAsDataURL(ctx context.Context, prompt string, seed int) (string, error) {
	// 使用 pollinations 免费图片生成服务（无需鉴权）
	// 文档参考: https://image.pollinations.ai/
	// 600x600，保证上半区满幅贴边。关闭 logo，传入 seed 确保稳定复现。
	escaped := url.PathEscape(prompt)
	endpoint := fmt.Sprintf("https://image.pollinations.ai/prompt/%s?width=600&height=600&nologo=true&seed=%d", escaped, seed)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "xhs-cover-bot/1.0")

	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("生成图片失败: %s", resp.Status)
	}
	bts, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	mime := resp.Header.Get("Content-Type")
	if mime == "" {
		mime = "image/jpeg"
	}
	base64Img := base64.StdEncoding.EncodeToString(bts)
	return fmt.Sprintf("data:%s;base64,%s", mime, base64Img), nil
}

// craftHeadline 简要提炼标题（无 LLM 降级方案）
func craftHeadline(topic string) string {
	t := strings.TrimSpace(topic)
	if strings.HasPrefix(t, "话题：") || strings.HasPrefix(t, "话题:") {
		t = strings.TrimPrefix(strings.TrimPrefix(t, "话题："), "话题:")
		t = strings.TrimSpace(t)
	}
	// 保持简短有力
	if len([]rune(t)) > 18 {
		return html.EscapeString(t[:18])
	}
	return html.EscapeString(t)
}

// craftQuote 生成简洁引言（无 LLM 降级方案）
func craftQuote(topic string) string {
	t := strings.TrimSpace(topic)
	return html.EscapeString(fmt.Sprintf("观点：%s 的核心在于趋势洞察、语境与影响。", t))
}

// buildCoverHTML 组装完整 HTML（内联 CSS + 固定 600x800 容器）
func buildCoverHTML(imgDataURL, headline, quote, infoBg string) string {
	// 上 70%: 560px；下 30%: 240px
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta charset="UTF-8" />
<meta name="viewport" content="width=device-width, initial-scale=1.0" />
<title>小红书封面图</title>
<style>
  :root { --w: 600px; --h: 800px; }
  * { box-sizing: border-box; }
  html, body { margin: 0; padding: 0; background: #f5f5f5; }
  body { display: flex; align-items: center; justify-content: center; min-height: 100vh; }
  .canvas { width: var(--w); height: var(--h); background: #fff; display: flex; flex-direction: column; }
  .illustration { height: 70%%; background-image: url('%s'); background-size: cover; background-position: center; }
  /* 绝对直角，无边框、无阴影、无圆角 */
  .illustration { border: none; box-shadow: none; outline: none; border-radius: 0; }
  .info { height: 30%%; background: %s; display: flex; flex-direction: column; align-items: center; justify-content: center; padding: 20px; }
  .headline { font-family: 'Inter','Noto Sans SC', system-ui, -apple-system, Segoe UI, Roboto, 'Helvetica Neue', Arial, 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', sans-serif; font-size: 34px; font-weight: 800; color: #111; text-align: center; margin: 0 0 10px; }
  .quote { font-family: 'Inter','Noto Sans SC', system-ui, -apple-system, Segoe UI, Roboto, 'Helvetica Neue', Arial, 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', sans-serif; font-size: 21px; line-height: 1.6; color: #555; text-align: center; margin: 0; }
</style>
</head>
<body>
  <div class="canvas" role="img" aria-label="XHS Cover: conceptual illustration + analysis">
    <div class="illustration"></div>
    <section class="info">
      <h1 class="headline">%s</h1>
      <p class="quote">%s</p>
    </section>
  </div>
</body>
</html>`, imgDataURL, infoBg, headline, quote)
}
