package services

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"regexp"
	"reverse-proxy/internal/core/ports"
)

const QueryRegex = `<(?:a|abbr|acronym|address|applet|area|audioscope|b|base|basefront|bdo|bgsound|big|blackface|blink|blockquote|body|bq|br|button|caption|center|cite|code|col|colgroup|comment|dd|del|dfn|dir|div|dl|dt|em|embed|fieldset|fn|font|form|frame|frameset|h1|head|hr|html|i|iframe|ilayer|img|input|ins|isindex|kdb|keygen|label|layer|legend|li|limittext|link|listing|map|marquee|menu|meta|multicol|nobr|noembed|noframes|noscript|nosmartquotes|object|ol|optgroup|option|p|param|plaintext|pre|q|rt|ruby|s|samp|script|select|server|shadow|sidebar|small|spacer|span|strike|strong|style|sub|sup|table|tbody|td|textarea|tfoot|th|thead|title|tr|tt|u|ul|var|wbr|xml|xmp)\\W`

var _ ports.IMiddlewareService = (*MiddlewareService)(nil)

type MiddlewareService struct {
	repo ports.IRepo
}

func NewMiddlewareService(repo ports.IRepo) ports.IMiddlewareService {
	return &MiddlewareService{
		repo: repo,
	}
}

func (m *MiddlewareService) BlockIp(c *fiber.Ctx) error {
	err := m.repo.StartConnection()
	if err != nil {
		fmt.Print(err)
		return err
	}
	blockedIPs, err := m.repo.GetAllBlockedIPs()
	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusForbidden)
	}

	clientIP := c.IP()
	for _, blockedIP := range blockedIPs {
		if clientIP == blockedIP {
			return c.SendStatus(fiber.StatusForbidden)
		}
	}

	return c.Next()
}

func (m *MiddlewareService) RewriteURIMiddleware(c *fiber.Ctx) error {
	host := c.Hostname()

	if len(host) >= 4 && host[:4] == "www." {
		newURI := "/site/www" + c.OriginalURL()
		c.Path(newURI)
	}

	return c.Next()
}

func (m *MiddlewareService) BlockRequestMiddleware(c *fiber.Ctx) error {
	blockRegex := regexp.MustCompile(QueryRegex)

	queryString := string(c.Request().URI().QueryString())
	if blockRegex.MatchString(queryString) {
		return c.SendStatus(fiber.StatusForbidden)
	}

	reqBody := c.Body()
	if blockRegex.Match(reqBody) {
		return c.SendStatus(fiber.StatusForbidden)
	}

	return c.Next()
}
