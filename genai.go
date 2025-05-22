package godo

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	genAIBasePath = "/v2/gen-ai/agents"
)

type GenAiService interface {
	CreateFunctionRoute(context.Context, string, *FunctionRouteCreateRequest) (*Agent, *Response, error)
	DeleteFunctionRoute(context.Context, string, string) (*Agent, *Response, error)
	UpdateFunctionRoute(context.Context, string, string, *FunctionRouteUpdateRequest) (*Agent, *Response, error)
}

type GenAiServiceOp struct {
	client *Client
}

type ChatbotIdentifiers struct {
	AgentChatbotIdentifier string `json:"agent_chatbot_identifier,omitempty"`
}
type Agent struct {
	AnthropicApiKey    *Info                `json:"anthropic_api_key,omitempty"`
	ApiKeyInfos        *Info                `json:"api_key_infos,omitempty"`
	ApiKeys            []*ApiKeys           `json:"api_keys,omitempty"`
	ChatBot            *ChatBot             `json:"chatbot,omitempty"`
	ChatbotIdentifiers []ChatbotIdentifiers `json:"chatbot_identifiers,omitempty"`
	CreatedAt          string               `json:"created_at,omitempty"`
	Deployment         *AgentDeployment     `json:"deployment,omitempty"`
	Descripton         string               `json:"description,omitempty"`
	UpdateAt           string               `json:"updated_at,omitempty"`
	Functions          []*Functions         `json:"functions,omitempty"`
	Guardrails         []*Guardrails        `json:"guardrails,omitempty"`
	IfCase             string               `json:"if_case,omitempty"`
	Instruction        string               `json:"instruction,omitempty"`
	K                  int                  `json:"k,omitempty"`
	KnowledgeBases     []*KnowledgeBase     `json:"knowledge_bases,omitempty"`
	MaxToken           int                  `json:"max_tokens,omitempty"`
	Model              *Model               `json:"model,omitempty"`
	Name               string               `json:"name,omitempty"`
	OpenAiApiKey       *OpenAiApiKey        `json:"open_ai_api_key,omitempty"`
	ProjectId          string               `json:"project_id,omitempty"`
	Region             string               `json:"region,omitempty"`
	RetrievalMethod    string               `json:"retrieval_method,omitempty"`
	RouteCreatedAt     string               `json:"route_created_at,omitempty"`
	RouteCreatedBy     string               `json:"route_created_by,omitempty"`
	RouteUuid          string               `json:"route_uuid,omitempty"`
	RouteName          string               `json:"route_name,omitempty"`
	Tags               []string             `json:"tags,omitempty"`
	Template           *AgentTemplate       `json:"template,omitempty"`
	Temperature        float64              `json:"temperature,omitempty"`
	TopP               float64              `json:"top_p,omitempty"`
	Url                string               `json:"url,omitempty"`
	UserId             string               `json:"user_id,omitempty"`
	Uuid               string               `json:"uuid,omitempty"`
}
type Agents struct {
	ChatBot            *ChatBot             `json:"chatbot,omitempty"`
	ChatbotIdentifiers []ChatbotIdentifiers `json:"chatbot_identifiers,omitempty"`
	Name               string               `json:"name,omitempty"`
	CreatedAt          time.Time            `json:"created_at,omitempty"`
	UpdateAt           time.Time            `json:"updated_at,omitempty"`
	Instruction        string               `json:"instruction,omitempty"`
	Descripton         string               `json:"description,omitempty"`
	IfCase             string               `json:"if_case,omitempty"`
	K                  int                  `json:"k,omitempty"`
	MaxToken           int                  `json:"max_tokens,omitempty"`
	ProjectId          string               `json:"project_id,omitempty"`
	Region             string               `json:"region,omitempty"`
	RetrievalMethod    string               `json:"retrieval_method,omitempty"`
	RouteCreatedAt     time.Time            `json:"route_created_at,omitempty"`
	RouteCreatedBy     string               `json:"route_created_by,omitempty"`
	RouteUuid          string               `json:"route_uuid,omitempty"`
	RouteName          string               `json:"route_name,omitempty"`
	Model              *Model               `json:"model,omitempty"`
	Deployment         *AgentDeployment     `json:"deployment,omitempty"`
	Tags               []string             `json:"tags,omitempty"`
	Template           *AgentTemplate       `json:"template,omitempty"`
	Temperature        float64              `json:"temperature,omitempty"`
	TopP               float64              `json:"top_p,omitempty"`
	Url                string               `json:"url,omitempty"`
	UserId             string               `json:"user_id,omitempty"`
	Uuid               string               `json:"uuid,omitempty"`
}

type genAIAgentsRoot struct {
	Agents []*Agents `json:"agents"`
	Links  *Links    `json:"links"`
	Meta   *Meta     `json:"meta"`
}
type AttachedGuardRails struct {
	IsDeleted bool   `json:"is_deleted,omitempty"`
	Name      string `json:"name,omitempty"`
	Priority  string `json:"priority,omitempty"`
	Uuid      string `json:"uuid,omitempty"`
}

type Guardrails struct {
	AgentUuid       string `json:"agent_uuid,omitempty"`
	CreatedAt       string `json:"created_at,omitempty"`
	DefaultResponse string `json:"default_response,omitempty"`
	Description     string `json:"description,omitempty"`
	GuardrailUuid   string `json:"guardrail_uuid,omitempty"`
	IsAttached      bool   `json:"is_attached,omitempty"`
	IsDefault       bool   `json:"is_default,omitempty"`
	Name            string `json:"name,omitempty"`
	Priority        int    `json:"priority,omitempty"`
	Type            string `json:"type,omitempty"`
	UpdatedAt       string `json:"updated_at,omitempty"`
	Uuid            string `json:"uuid,omitempty"`
}
type AttachedKnowledgebases struct {
	IsDeleted bool   `json:"is_deleted,omitempty"`
	Name      string `json:"name,omitempty"`
	Uuid      string `json:"uuid,omitempty"`
}

type ApiKeys struct {
	ApiKey string `json:"api_key,omitempty"`
}

type Info struct {
	CreatedAt string `json:"created_at,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	DeletedAt string `json:"deleted_at,omitempty"`
	Name      string `json:"name,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Uuid      string `json:"uuid,omitempty"`
}
type Model struct {
	Agreement        *Agreement `json:"agreement,omitempty"`
	CreatedAt        string     `json:"created_at,omitempty"`
	InferenceName    string     `json:"inference_name,omitempty"`
	InferenceVersion string     `json:"inference_version,omitempty"`
	IsFoundational   bool       `json:"is_foundational,omitempty"`
	Name             string     `json:"name,omitempty"`
	ParentUuid       string     `json:"parent_uuid,omitempty"`
	Provider         string     `json:"provider,omitempty"`
	UpdatedAt        string     `json:"updated_at,omitempty"`
	UploadComplete   bool       `json:"upload_complete,omitempty"`
	Url              string     `json:"url,omitempty"`
	Usecases         []string   `json:"usecases,omitempty"`
	Uuid             string     `json:"uuid,omitempty"`
	Version          *Version   `json:"version,omitempty"`
}

type Agreement struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	Url         string `json:"url,omitempty"`
	Uuid        string `json:"uuid,omitempty"`
}

type Version struct {
	Major int `json:"major,omitempty"`
	Minor int `json:"minor,omitempty"`
	Patch int `json:"patch,omitempty"`
}

type OpenAiApiKey struct {
	CreatedAt string   `json:"created_at,omitempty"`
	CreatedBy string   `json:"created_by,omitempty"`
	DeletedAt string   `json:"deleted_at,omitempty"`
	Models    []string `json:"models,omitempty"`
	Name      string   `json:"name,omitempty"`
	UpdatedAt string   `json:"updated_at,omitempty"`
	Uuid      string   `json:"uuid,omitempty"`
}
type KnowledgeBase struct {
	AddedToAgentAt     string           `json:"added_to_agent_at,omitempty"`
	CreatedAt          string           `json:"created_at,omitempty"`
	DatabaseId         string           `json:"database_id,omitempty"`
	EmbeddingModelUuid string           `json:"embedding_model_uuid,omitempty"`
	IsPublic           bool             `json:"is_public,omitempty"`
	LastIndexingJob    *LastIndexingJob `json:"last_indexing_job,omitempty"`
	Name               string           `json:"name,omitempty"`
	ProjectId          string           `json:"project_id,omitempty"`
	Region             string           `json:"region,omitempty"`
	Tags               []string         `json:"tags,omitempty"`
	UpdateAt           string           `json:"updated_at,omitempty"`
	UserId             string           `json:"user_id,omitempty"`
	Uuid               string           `json:"uuid,omitempty"`
}
type LastIndexingJob struct {
	CompletedDatasources int      `json:"completed_datasources,omitempty"`
	CreatedAt            string   `json:"created_at,omitempty"`
	DataSourceUuids      []string `json:"data_source_uuids,omitempty"`
	FinishedAt           string   `json:"finished_at,omitempty"`
	KnowledgeBaseUuid    string   `json:"knowledge_base_uuid,omitempty"`
	Phase                string   `json:"phase,omitempty"`
	StartedAt            string   `json:"started_at,omitempty"`
	Tokens               int      `json:"tokens,omitempty"`
	TotalDatasources     int      `json:"total_datasources,omitempty"`
	UpdatedAt            string   `json:"updated_at,omitempty"`
	Uuid                 string   `json:"uuid,omitempty"`
}
type AgentDeployment struct {
	CreatedAt  string `json:"created_at,omitempty"`
	Name       string `json:"name,omitempty"`
	Status     string `json:"status,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	Url        string `json:"url,omitempty"`
	Uuid       string `json:"uuid,omitempty"`
	Visibility string `json:"visibility,omitempty"`
}

type ChatBot struct {
	ButtonBackgroundColor string `json:"button_background_color,omitempty"`
	Logo                  string `json:"logo,omitempty"`
	Name                  string `json:"name,omitempty"`
	PrimaryColor          string `json:"primary_color,omitempty"`
	SecondaryColor        string `json:"secondary_color,omitempty"`
	StartingMessage       string `json:"starting_message,omitempty"`
}
type AgentVisibilityUpdateRequest struct {
	Uuid       string `json:"uuid,omitempty"`
	Visibility string `json:"visibility,omitempty"`
}

type AgentTemplate struct {
	CreatedAt      string           `json:"created_at,omitempty"`
	Instruction    string           `json:"instruction,omitempty"`
	Descripton     string           `json:"description,omitempty"`
	K              int              `json:"k,omitempty"`
	KnowledgeBases []*KnowledgeBase `json:"knowledge_bases,omitempty"`
	MaxToken       int              `json:"max_tokens,omitempty"`
	Model          *Model           `json:"model,omitempty"`
	Name           string           `json:"name,omitempty"`
	Temperature    float64          `json:"temperature,omitempty"`
	TopP           float64          `json:"top_p,omitempty"`
	UpdateAt       string           `json:"updated_at,omitempty"`
	Uuid           string           `json:"uuid,omitempty"`
}
type FunctionRouteCreateRequest struct {
	AgentUuid     string `json:"agent_uuid,omitempty"`
	Description   string `json:"description,omitempty"`
	FaasName      string `json:"faas_name,omitempty"`
	FaasNamespace string `json:"faas_namespace,omitempty"`
	FunctionName  string `json:"function_name,omitempty"`
	InputSchema   string `json:"input_schema,omitempty"`
	OutputSchema  string `json:"output_schema,omitempty"`
}

type FunctionRouteUpdateRequest struct {
	AnthropicKeyUuid string   `json:"anthropic_key_uuid,omitempty"`
	Description      string   `json:"description,omitempty"`
	Instruction      string   `json:"instruction,omitempty"`
	K                int      `json:"k,omitempty"`
	MaxToken         int      `json:"max_tokens,omitempty"`
	ModelUuid        string   `json:"model_uuid,omitempty"`
	Name             string   `json:"name,omitempty"`
	OpenAiKeyUuid    string   `json:"open_ai_key_uuid,omitempty"`
	ProjectId        string   `json:"project_id,omitempty"`
	Region           string   `json:"region,omitempty"`
	Tags             []string `json:"tags,omitempty"`
	Temperature      float64  `json:"temperature,omitempty"`
	TopP             float64  `json:"top_p,omitempty"`
	Uuid             string   `json:"uuid,omitempty"`
}
type genAIAgentRoot struct {
	Agent *Agent `json:"agent"`
}

// CreateFunctionRoute implements GenAiService.
func (g *GenAiServiceOp) CreateFunctionRoute(ctx context.Context, id string, create *FunctionRouteCreateRequest) (*Agent, *Response, error) {
	path := fmt.Sprintf("%s/%s/functions", genAIBasePath, id)
	req, err := g.client.NewRequest(ctx, http.MethodPost, path, create)
	if err != nil {
		return nil, nil, err
	}

	root := new(genAIAgentRoot)
	resp, err := g.client.Do(ctx, req, root)

	if err != nil {
		return nil, resp, err
	}

	return root.Agent, resp, nil
}

// DeleteFunctionRoute implements GenAiService.
func (g *GenAiServiceOp) DeleteFunctionRoute(ctx context.Context, agent_id string, function_id string) (*Agent, *Response, error) {
	path := fmt.Sprintf("%s/%s/functions/%s", genAIBasePath, agent_id, function_id)
	req, err := g.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(genAIAgentRoot)
	resp, err := g.client.Do(ctx, req, root)

	if err != nil {
		return nil, resp, err
	}

	return root.Agent, resp, nil
}

// UpdateFunctionRoute implements GenAiService.
func (g *GenAiServiceOp) UpdateFunctionRoute(ctx context.Context, agent_id string, function_id string, update *FunctionRouteUpdateRequest) (*Agent, *Response, error) {
	path := fmt.Sprintf("%s/%s/functions/%s", genAIBasePath, agent_id, function_id)
	req, err := g.client.NewRequest(ctx, http.MethodPut, path, update)
	if err != nil {
		return nil, nil, err
	}

	root := new(genAIAgentRoot)
	resp, err := g.client.Do(ctx, req, root)

	if err != nil {
		return nil, resp, err
	}

	return root.Agent, resp, nil
}

type Functions struct {
	ApiKey        string `json:"api_key,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	Description   string `json:"description,omitempty"`
	GuardrailUuid string `json:"guardrail_uuid,omitempty"`
	FaasName      bool   `json:"faas_name,omitempty"`
	FaasNamespace bool   `json:"faas_namespace,omitempty"`
	Name          string `json:"name,omitempty"`
	Type          string `json:"type,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
	Url           string `json:"url,omitempty"`
	Uuid          string `json:"uuid,omitempty"`
}
